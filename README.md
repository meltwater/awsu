# Assume AWS IAM role
[![Travis CI](https://img.shields.io/travis/meltwater/awsu/master.svg)](https://travis-ci.org/meltwater/awsu)

Assumes an IAM role and passes the temporary credentials to another command or shell.

If you manage multiple AWS accounts and use IAM role switching to perform work in them, this would
allow you to use tools like Terraform, Docker Machine or Vagrant in the accounts. Cross account
IAM role switching is described at

* https://aws.amazon.com/blogs/aws/new-cross-account-access-in-the-aws-management-console/

## Usage

```
Assume a AWS IAM role and execute a command or shell. If no command is given an interactive
shell will be started with the credentials supplied as environment variables.

Usage:
  awsu IAMRoleARN [command] [args]... [flags]

Flags:
      --duration int   Expiration time in seconds for the temporary credentials (default 900)
```

## Installation
See the [releases page](https://github.com/meltwater/awsu/releases) for version numbers

```
SECRETARY_VERSION=x.y.z

sudo curl -fsSLo /usr/bin/awsu "https://github.com/meltwater/awsu/releases/download/${SECRETARY_VERSION}/awsu-`uname -s`-`uname -m`" && \
     chmod +x /usr/bin/awsu
```

If you have Golang installed you could also install from source like

```
go get github.com/meltwater/awsu
```

### Quick Access
It could be useful to setup ~/.bash_aliases for roles in different accounts

```
# Enable prompt postfix
if [[ -n "$PROMPT_POSTFIX" ]]; then
    PS1="${PS1}${PROMPT_POSTFIX}\$ "
fi

# Easy color customization at http://ezprompt.net/
alias ondev='env PROMPT_POSTFIX="\[\e[1;32m\]<dev>\[\e[m\]" awsu arn:aws:iam::123456789:role/Developer'
alias onstaging='env PROMPT_POSTFIX="\[\e[1;33m\]<staging>\[\e[m\]" awsu arn:aws:iam::678912345:role/Developer'
alias onproduction='env PROMPT_POSTFIX="\[\e[1;31m\]<production>\[\e[m\]" awsu arn:aws:iam::891234567:role/Developer'
```

For example

```
$ ondev terraform plan
...

$ ondev docker-machine create --driver amazonec2 ...
...
```

Or to drop into an interactive shell

```
$ onstaging
<staging>$ terraform plan
...

<staging>$ terraform apply
...
```

Inspired by

* https://github.com/mlrobinson/aws-profile
* https://github.com/jbuck/assume-aws-role
* http://blog.sinica.me/aws_multi_account_with_terraform.html

