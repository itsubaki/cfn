package stack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group name) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("config"))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cf.New(session.Must(session.NewSession()))
	for _, tmplpath := range config.Template() {
		fmt.Print(tmplpath)

		group := c.Args().Get(0)
		name := cfg.StackName(group, tmplpath)
		body, err := cfg.TemplateBody(tmplpath)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		key := "StackGroupName"
		tags := append(config.Tag(), &cf.Tag{Key: &key, Value: &group})

		iam := "CAPABILITY_IAM"
		niam := "CAPABILITY_NAMED_IAM"
		req := &cf.CreateStackInput{
			StackName:    &name,
			TemplateBody: &body,
			Capabilities: []*string{&iam, &niam},
			Tags:         tags,
		}

		res, err := client.CreateStack(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackCreateComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println(" created. " + *res.StackId)
	}
}
