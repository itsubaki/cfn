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
	list := config.Template()
	for i := 0; i < len(list); i++ {
		tmplpath := list[i]
		fmt.Print(tmplpath)

		body, err := cfg.TemplateBody(tmplpath)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		req := &cf.EstimateTemplateCostInput{TemplateBody: &body}
		res, err := client.EstimateTemplateCost(req)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println(" " + *res.Url)
	}
}
