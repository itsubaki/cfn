package stack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Update(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	opts := session.Options{SharedConfigState: session.SharedConfigEnable}
	client := cf.New(session.Must(session.NewSessionWithOptions(opts)))
	for _, template := range config.Resources {
		group := c.Args().Get(0)
		name := cfg.StackName(group, template.Name)
		fmt.Print(name)

		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		iam := "CAPABILITY_IAM"
		niam := "CAPABILITY_NAMED_IAM"
		create := &cf.UpdateStackInput{
			StackName:    &name,
			TemplateBody: &body,
			Capabilities: []*string{&iam, &niam},
			Parameters:   template.Parameter(),
			Tags:         config.Tag(),
		}

		res, err := client.UpdateStack(create)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		desc := &cf.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackUpdateComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" updated. " + *res.StackId)
	}
}
