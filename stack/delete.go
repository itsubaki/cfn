package stack

import (
	"fmt"
	"os"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	ses "github.com/itsubaki/cfn/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Delete(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	for _, template := range config.Reverse() {
		group := c.Args().Get(0)
		name := cfg.StackName(group, template.Name)
		fmt.Print(name)

		client := cf.New(ses.New(template.Region))
		req := &cf.DeleteStackInput{StackName: &name}
		_, err := client.DeleteStack(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		desc := &cf.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackDeleteComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" deleted.")
	}
}
