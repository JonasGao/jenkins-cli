# jenkins-cli

```
NAME:
    jk - Make Jenkins like a nice JK

USAGE:
    jk [global options] command [command options] [arguments...]

COMMANDS:
get, g      Get the job
   build, b    Build the job
   history, h  Get the job history builds
   latest, lt  Get the job latest build
   list, l     List all jobs
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
--help, -h  show help (default: false)
```

## Usage

First create ".jk.yaml" config file in "$HOME".

```yaml
domain: https://jenkins.com
username: admin
password: 123456
```
