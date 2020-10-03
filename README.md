# cfn

 - cfn is a command line tool for AWS CloudFormation.
 - cfn execute multiple template at once.

## Install

```console
$ go get github.com/itsubaki/cfn
```

## Example

```console
$ ls
example.yaml	template
$ cat example.yaml
resources:
  - name: vpc
    type: template/vpc.yml
    properties:
      - name: ProjectName
        value: test
      - name: Environment
        value: develop
  - name: subnet
    type: template/subnet.yml
  - name: sg
    type: tempalte/sg.yml

tags:
  - name: StackGroup
    value: test

// stack name is ${stack_group_name}-{resource.name}
$ cfn stack create test
template/vpc.yaml    created. test-vpc
template/subnet.yaml created. test-subnet
template/sg.yaml     created. test-sg

$ cfn changeset create test
template/vpc.yaml    no update.
template/subnet.yaml no update.
template/sg.yaml     created. changeset-test-sg-1497174211

$ cfn changeset describe changeset-test-sg-1497174211
{
  ChangeSetId: "arn:aws:cloudformation:ap-northeast-1:************:changeSet/changeset-test-sg-1497174211/********-****-****-****-************",
  ChangeSetName: "changeset-test-sg-1497174211",
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
  StackId: "arn:aws:cloudformation:ap-northeast-1:************:stack/test-sg/********-****-****-****-************",
  StackName: "test-sg",
  Status: "CREATE_COMPLETE",
  Tags: [{
      Key: "ProjectName",
      Value: "test"
    }]
}

$ cfn changeset execute changeset-test-sg-1497174211
test-sg updated. changeset-test-sg-1497174211
```
