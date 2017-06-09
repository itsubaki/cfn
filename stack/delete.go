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

func Delete(c *cli.Context) {
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
	for i := len(list); i > 0; i-- {
		tmplname := list[i-1].(string)
		fmt.Println(tmplname)

		tmp := strings.Replace(tmplname, "/", "-", -1)
		suffix := strings.Replace(tmp, ".yaml", "", -1)
		name := c.Args().Get(0) + "-" + suffix

		delete := &cf.DeleteStackInput{
			StackName: &name,
		}

		_, err := client.DeleteStack(delete)
		if err != nil {
			fmt.Println(err)
			continue
		}

		desc := &cf.DescribeStacksInput{
			StackName: &name,
		}
		err = client.WaitUntilStackDeleteComplete(desc)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}
