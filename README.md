# cfn

 - cfn is command line tool for AWS CloudFormation.
 - cfn execute multiple template at once.

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

 $ cfn stack create test
 template/vpc.yaml created.
 template/subnet.yaml created.
 template/sg.yaml created.

 $ cfn stack delete test
 template/sg.yaml deleted.
 template/subnet.yaml deleted.
 template/vpc.yaml deleted.
 ```
