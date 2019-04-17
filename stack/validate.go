package stack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"

	cli "gopkg.in/urfave/cli.v1"
)

func Validate(c *cli.Context) {
	config, err := cfg.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	opts := session.Options{SharedConfigState: session.SharedConfigEnable}
	client := cf.New(session.Must(session.NewSessionWithOptions(opts)))
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
