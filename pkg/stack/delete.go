package stack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/itsubaki/cfn/pkg/config"
	"github.com/itsubaki/cfn/pkg/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Delete(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	conf, err := config.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	for _, template := range conf.Reverse() {
		group := c.Args().Get(0)
		name := config.StackName(group, template.Name)
		fmt.Print(name)

		client := cloudformation.New(session.New(template.Region))
		req := &cloudformation.DeleteStackInput{StackName: &name}
		_, err := client.DeleteStack(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		desc := &cloudformation.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackDeleteComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" deleted.")
	}
}
