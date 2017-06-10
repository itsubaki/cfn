package stack

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	cli "gopkg.in/urfave/cli.v1"
	yaml "gopkg.in/yaml.v2"
)

func Create(c *cli.Context) {
	if len(c.Args()) < 1 {
		fmt.Println("error: stack name is null.")
		os.Exit(1)
	}

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
		tmplpath := list[i].(string)
		fmt.Print(tmplpath)

		buf, err := ioutil.ReadFile(tmplpath)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		tmp := strings.Replace(tmplpath, "/", "-", -1)
		suffix := strings.Replace(tmp, ".yaml", "", -1)
		name := c.Args().Get(0) + "-" + suffix

		body := string(buf)
		create := &cf.CreateStackInput{
			StackName:    &name,
			TemplateBody: &body,
		}

		res, err := client.CreateStack(create)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeStacksInput{StackName: &name}
		err = client.WaitUntilStackCreateComplete(desc)
		if err != nil {
			fmt.Println()
			fmt.Println(err)
			continue
		}

		fmt.Println(" created. " + *res.StackId)
	}
}
