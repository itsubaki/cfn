package stack

import (
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Update(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: stack group name is null.")
		os.Exit(1)
	}

	config, err := cfg.Read(c.String("config"))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cf.New(session.Must(session.NewSession()))

	list := config["Templates"].([]interface{})
	for i := 0; i < len(list); i++ {
		tmpl := list[i].(string)
		fmt.Print(tmpl)

		tmp := strings.Replace(tmpl, "/", "-", -1)
		suffix := strings.Replace(tmp, ".yaml", "", -1)
		name := c.Args().Get(0) + "-" + suffix

		body, err := cfg.TemplateBody(tmpl)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		create := &cf.UpdateStackInput{
			StackName:    &name,
			TemplateBody: &body,
		}

		res, err := client.UpdateStack(create)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackUpdateComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println(" updated. " + *res.StackId)
	}
}
