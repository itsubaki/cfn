package stack

import (
	"fmt"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	ses "github.com/itsubaki/cfn/session"

	cli "gopkg.in/urfave/cli.v1"
)

func Validate(c *cli.Context) {
	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	client := cf.New(ses.New())
	for _, template := range config.Resources {
		fmt.Print(template.Name)

		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		req := &cf.ValidateTemplateInput{TemplateBody: &body}
		_, err = client.ValidateTemplate(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" ok.")
	}

}
