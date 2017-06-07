# cfn
AWS CloudFormation command line tool.

## cfn.yaml

```yaml
template:
  - template/vpc.yaml
  - template/sg.yaml
  - template/s3.yaml
  - template/elb.yaml
  - template/ec2.yaml
parameters:
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
     validate, v    Validates a specified template.
     stack, s       Create, Update, Delete, Describe Stack.
     changeset, cs  Create, Execute, Delete, Describe Changeset.
     help, h        Shows a list of commands or help for one command

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
     create, c       Creates a stack as specified in the template.
     update, u       Updates a stack as specified in the template.
     delete, d       Deletes a specified stack.
     describe, desc  Returns the description for the specified stack.

OPTIONS:
   --help, -h  show help
```

```console
$ cfn changeset -h
NAME:
   cfn changeset - Create, Execute, Delete, Describe Changeset.

USAGE:
   cfn changeset command [command options] [arguments...]

COMMANDS:
     create, c       Creates a list of changes that will be applied to a stack so that you can review the changes before executing them.
     execute, e      Updates a stack using the input information that was provided when the specified change set was created.
     delete, d       Deletes the specified change set.
     describe, desc  Returns the inputs for the change set and a list of changes that AWS CloudFormation will make if you execute the change set.

OPTIONS:
   --help, -h  show help
```
