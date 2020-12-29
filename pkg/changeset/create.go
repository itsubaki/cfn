package changeset

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/itsubaki/cfn/pkg/config"
	"github.com/itsubaki/cfn/pkg/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	conf, err := config.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	client := cloudformation.New(session.New())
	for _, template := range conf.Resources {
		fmt.Print(template.Name)

		name := config.StackName(c.Args().Get(0), template.Name)
		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		changeSetName := "changeset-" + name + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		iam := "CAPABILITY_IAM"
		niam := "CAPABILITY_NAMED_IAM"
		req := &cloudformation.CreateChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
			TemplateBody:  &body,
			Capabilities:  []*string{&iam, &niam},
			Parameters:    template.Parameter(),
			Tags:          conf.Tag(),
		}

		res, err := client.CreateChangeSet(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		desc := &cloudformation.DescribeChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
		}
		err = client.WaitUntilChangeSetCreateComplete(desc)
		if err == nil {
			fmt.Println(" created. " + *res.Id)
			break
		}

		// Delete ChangeSet of Failed Status
		input := &cloudformation.DeleteChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
		}
		_, err = client.DeleteChangeSet(input)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" no update.")
	}
}
