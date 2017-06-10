package stack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"

	cli "gopkg.in/urfave/cli.v1"
)

func Validate(c *cli.Context) {
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

		body, err := cfg.TemplateBody(tmpl)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		input := &cf.ValidateTemplateInput{TemplateBody: &body}
		_, err = client.ValidateTemplate(input)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println(" ok.")
	}

}
