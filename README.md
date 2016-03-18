# Assume AWS IAM role
[![Travis CI](https://img.shields.io/travis/meltwater/awsu/master.svg)](https://travis-ci.org/meltwater/awsu)

Assumes an IAM role and passes the temporary credentials to another command or shell.

If you manage multiple AWS accounts and use IAM role switching to perform work in them, this would
allow you to use tools like Terraform, Docker Machine or Vagrant in the accounts. Cross account
IAM role switching is described at

* https://aws.amazon.com/blogs/aws/new-cross-account-access-in-the-aws-management-console/

## Usage

```
Assume a AWS IAM role and execute a command or shell

Usage:
  awsu IAMRoleARN [command] [args]... [flags]

Flags:
      --duration int   Expiration time in seconds for the temporary credentials (default 900)
```

It could be useful to setup ~/.bash_aliases for roles in different accounts

```
alias ondev='awsu arn:aws:iam::123456789:role/Developer'
alias onprod='awsu arn:aws:iam::891234567:role/Developer'
```

For example

```
$ ondev terraform plan
$ ondev docker-machine create --driver amazonec2 ...
```

Inspired by

* https://github.com/mlrobinson/aws-profile
* https://github.com/jbuck/assume-aws-role
* http://blog.sinica.me/aws_multi_account_with_terraform.html
