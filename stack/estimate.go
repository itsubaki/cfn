package stack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cfg "github.com/itsubaki/cfn/config"
	cli "gopkg.in/urfave/cli.v1"
)

func Estimate(c *cli.Context) {

	config, err := cfg.Read(c.String("config"))
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cf.New(session.Must(session.NewSession()))
	for _, template := range config.Resources {
		fmt.Print(template.Name)

		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		req := &cf.EstimateTemplateCostInput{TemplateBody: &body}
		res, err := client.EstimateTemplateCost(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" " + *res.Url)
	}
}
