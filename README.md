# cfn

 - cfn is command line tool for AWS CloudFormation.
 - cfn execute multiple template at once.

## Install

```console
$ go get github.com/itsubaki/cfn
```


## Example

```console
$ ls
cfn.yaml	template
$ cat cfn.yaml
Templates:
 - template/vpc.yaml
 - template/subnet.yaml
 - template/sg.yaml
Parameters:
   - ProjectName: test
Tags:
   - ProjectName: test

// first argument is stack group name
// stack name is ${stack_group_name}-{template_dir}-{template_name} without file extension
$ cfn stack create test
template/vpc.yaml    created. test-template-vpc
template/subnet.yaml created. test-template-subnet
template/sg.yaml     created. test-template-sg

$ cfn changeset create test
template/sg.yaml created. changeset-test-template-sg-1497174211

$ cfn changeset execute changeset-test-template-sg-1497174211
test-template-sg updated. changeset-test-template-sg-1497174211
```
