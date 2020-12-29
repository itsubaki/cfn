package stack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/itsubaki/cfn/pkg/config"
	"github.com/itsubaki/cfn/pkg/session"
	cli "gopkg.in/urfave/cli.v1"
)

func Estimate(c *cli.Context) {
	conf, err := config.Read(c.String("file"))
	if err != nil {
		fmt.Printf("read file: %v\n", err)
		return
	}

	client := cloudformation.New(session.New())
	for _, template := range conf.Resources {
		fmt.Print(template.Name)

		body, err := template.Body()
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		req := &cloudformation.EstimateTemplateCostInput{TemplateBody: &body}
		res, err := client.EstimateTemplateCost(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			break
		}

		fmt.Println(" " + *res.Url)
	}
}
