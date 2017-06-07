# cfn
AWS CloudFormation command line tool.

## credential.yaml

```yaml
access_key:
secret_key:
```

## config.yaml

```yaml
template:
  - template/vpc.yaml
  - template/sg.yaml
  - template/s3.yaml
  - template/elb.yaml
  - template/ec2.yaml
vars:
  ProjectName: myproject
  KeyPair: mykeypair
  CertificateArn:
```

## Usage

```console
$ cfn -h
NAME:
   cfn

USAGE:
   cfn [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
     validate, v  Validates a specified template.
     stack, s     Create, Update, Delete, Describe Stack.
     help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

```console
$ cfn stack -h
NAME:
   cfn stack - Create, Update, Delete, Describe Stack.

USAGE:
   cfn stack command [command options] [arguments...]

COMMANDS:
     create    Creates a stack as specified in the template.
     update    Updates a stack as specified in the template.
     delete    Deletes a specified stack.
     describe  Returns the description for the specified stack.

OPTIONS:
   --help, -h  show help
```
