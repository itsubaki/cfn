package changeset

import (
	"fmt"
	"os"
	"strconv"
	"time"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	ses "github.com/itsubaki/cfn/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	client := cf.New(ses.New())
	for _, template := range config.Resources {
		fmt.Print(template.Name)

		name := cfg.StackName(c.Args().Get(0), template.Name)
		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		changeSetName := "changeset-" + name + "-" + strconv.FormatInt(time.Now().Unix(), 10)
		iam := "CAPABILITY_IAM"
		niam := "CAPABILITY_NAMED_IAM"
		req := &cf.CreateChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
			TemplateBody:  &body,
			Capabilities:  []*string{&iam, &niam},
			Parameters:    template.Parameter(),
			Tags:          config.Tag(),
		}

		res, err := client.CreateChangeSet(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		desc := &cf.DescribeChangeSetInput{
			ChangeSetName: &changeSetName,
			StackName:     &name,
		}
		err = client.WaitUntilChangeSetCreateComplete(desc)
		if err == nil {
			fmt.Println(" created. " + *res.Id)
			break
		}

		// Delete ChangeSet of Failed Status
		input := &cf.DeleteChangeSetInput{
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
