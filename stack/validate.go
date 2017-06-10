package stack

import (
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"

	cli "gopkg.in/urfave/cli.v1"
	yaml "gopkg.in/yaml.v2"
)

func Validate(c *cli.Context) {
	cpath := c.String("config")

	buf, err := ioutil.ReadFile(cpath)
	if err != nil {
		fmt.Println(err)
		return
	}

	config := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := cf.New(session.Must(session.NewSession()))

	list := config["Templates"].([]interface{})
	for i := 0; i < len(list); i++ {
		tmplpath := list[i].(string)
		fmt.Print(tmplpath)

		buf, err := ioutil.ReadFile(tmplpath)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		body := string(buf)
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
