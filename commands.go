package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

func filterExistingCredentials(list []string) []string {
	var result []string
	for _, item := range list {
		if !strings.HasPrefix(item, "AWS_ACCESS_KEY_ID=") &&
			!strings.HasPrefix(item, "AWS_SECRET_ACCESS_KEY=") &&
			!strings.HasPrefix(item, "AWS_SESSION_TOKEN=") {
			result = append(result, item)
		}
	}

	return result
}

// Encrypts data from stdin and writes to stdout
func executeCommand(durationSeconds int64, iamRole string, args []string) {
	hostname, err := os.Hostname()
	sessionName := fmt.Sprintf("%s-%s-%s",
		defaults(os.Getenv("USER"), "unknown"),
		defaults(hostname, os.Getenv("HOST"), os.Getenv("HOSTNAME"), "unknown"),
		randSeq(8))

	svc := sts.New(session.New())

	params := &sts.AssumeRoleInput{
		RoleArn:         aws.String(iamRole),     // Required
		RoleSessionName: aws.String(sessionName), // Required
		DurationSeconds: aws.Int64(durationSeconds),
		//		ExternalId:      aws.String("externalIdType"),
		//		Policy:          aws.String("sessionPolicyDocumentType"),
		//		SerialNumber:    aws.String("serialNumberType"),
		//		TokenCode:       aws.String("tokenCodeType"),
	}

	resp, err := svc.AssumeRole(params)
	check(err)

	// Default to launch a subshell
	binary := defaults(os.Getenv("SHELL"), "/bin/sh")

	// Resolve absolute path of binary
	if len(args) >= 1 {
		binary, err = exec.LookPath(args[0])
		check(err)
	}

	// Inject the temporary credentials
	env := append(filterExistingCredentials(os.Environ()),
		fmt.Sprintf("AWS_ACCESS_KEY_ID=%s", *resp.Credentials.AccessKeyId),
		fmt.Sprintf("AWS_SECRET_ACCESS_KEY=%s", *resp.Credentials.SecretAccessKey),
		fmt.Sprintf("AWS_SESSION_TOKEN=%s", *resp.Credentials.SessionToken))

	// Execute subcommand
	err = syscall.Exec(binary, args, env)
	check(err)
}
