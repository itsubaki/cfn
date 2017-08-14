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
resources:
  - name: backend
    type: template/backend.yml
    parameters:
      - name: ProjectName
        value: test
      - name: Environment
        value: develop
tags:
  - name: StackGroup
    value: test

// first argument is stack group name
// stack name is ${stack_group_name}-{template_name} without file extension
$ cfn stack create test
template/vpc.yaml    created. test-template-vpc
template/subnet.yaml created. test-template-subnet
template/sg.yaml     created. test-template-sg

$ cfn changeset create test
template/vpc.yaml    no update.
template/subnet.yaml no update.
template/sg.yaml     created. changeset-test-template-sg-1497174211

$ cfn changeset describe changeset-test-template-sg-1497174211
{
  ChangeSetId: "arn:aws:cloudformation:ap-northeast-1:************:changeSet/changeset-test-template-sg-1497174211/********-****-****-****-************",
  ChangeSetName: "changeset-test-template-sg-1497174211",
  Changes: [{
      ResourceChange: {
        Action: "Modify",
        Details: [{
            ChangeSource: "DirectModification",
            Evaluation: "Static",
            Target: {
              Attribute: "Properties",
              Name: "SecurityGroupIngress",
              RequiresRecreation: "Never"
            }
          }],
        LogicalResourceId: "PublicSecurityGroup",
        PhysicalResourceId: "sg-d3b57fb5",
        Replacement: "False",
        ResourceType: "AWS::EC2::SecurityGroup",
        Scope: ["Properties"]
      },
      Type: "Resource"
    }],
  CreationTime: 2017-06-11 12:22:23.522 +0000 UTC,
  ExecutionStatus: "AVAILABLE",
  Parameters: [{
      ParameterKey: "ProjectName",
      ParameterValue: "test"
    }],
  StackId: "arn:aws:cloudformation:ap-northeast-1:************:stack/test-template-sg/********-****-****-****-************",
  StackName: "test-template-sg",
  Status: "CREATE_COMPLETE",
  Tags: [{
      Key: "ProjectName",
      Value: "test"
    }]
}

$ cfn changeset execute changeset-test-template-sg-1497174211
test-template-sg updated. changeset-test-template-sg-1497174211
```
