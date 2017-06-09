package stack

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cli "gopkg.in/urfave/cli.v1"
	yaml "gopkg.in/yaml.v2"
)

func Create(c *cli.Context) {
	path := c.String("config")

	buf, err := ioutil.ReadFile(path)
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
		tmplname := list[i].(string)
		fmt.Println(tmplname)

		buf, err := ioutil.ReadFile(tmplname)
		if err != nil {
			fmt.Println(err)
			continue
		}

		tmp := strings.Replace(tmplname, "/", "-", -1)
		suffix := strings.Replace(tmp, ".yaml", "", -1)
		name := c.Args().Get(0) + "-" + suffix

		tmpl := string(buf)
		create := &cf.CreateStackInput{
			StackName:    &name,
			TemplateBody: &tmpl,
		}

		out, err := client.CreateStack(create)
		if err != nil {
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeStacksInput{
			StackName: &name,
		}
		err = client.WaitUntilStackCreateComplete(desc)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println(out)
	}
}
