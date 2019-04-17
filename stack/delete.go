package stack

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Delete(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: first argument(stack group) is required")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Println("read file: %v", err)
		return
	}

	opts := session.Options{SharedConfigState: session.SharedConfigEnable}
	client := cf.New(session.Must(session.NewSessionWithOptions(opts)))
	for _, template := range config.Reverse() {
		group := c.Args().Get(0)
		name := cfg.StackName(group, template.Name)
		fmt.Print(name)

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

		fmt.Println(" deleted. " + name)
	}
}
