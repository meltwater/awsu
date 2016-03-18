package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-errors/errors"
	"github.com/spf13/cobra"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var durationSeconds int64

	rootCmd := &cobra.Command{
		Use:   "awsu IAMRoleARN [command] [args]...",
		Short: "Assume a AWS IAM role and execute a command or shell",
		Run: func(cmd *cobra.Command, args []string) {
			assertThat(len(args) >= 1, "Expected an IAM role")
			executeCommand(durationSeconds, args[0], args[1:])
		},
	}

	rootCmd.Flags().Int64VarP(&durationSeconds, "duration", "", int64(900), "Expiration time in seconds for the temporary credentials")

	// Handle checked errors nicely
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *CommandError:
				fmt.Fprintf(os.Stderr, "%s\n", err)
			default:
				fmt.Fprintf(os.Stderr, "%s\n", errors.Wrap(err, 2).ErrorStack())
			}

			os.Exit(1)
		}
	}()

	rootCmd.Execute()
}
