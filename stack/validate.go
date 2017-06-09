package stack

import (
	"fmt"
	"io/ioutil"

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

	list := config["Templates"].([]interface{})
	for i := 0; i < len(list); i++ {
		fmt.Println(list[i])
	}

}
