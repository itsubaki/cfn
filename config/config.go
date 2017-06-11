package config

import (
	"io/ioutil"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	yaml "gopkg.in/yaml.v2"
)

type Config map[string]interface{}
type TemplateList []string
type TagList []*cf.Tag

func Read(path string) (Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := make(Config)
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func TemplateBody(path string) (string, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(buf), nil
}

func (c Config) Template() TemplateList {
	template := TemplateList{}
	list := c["Templates"].([]interface{})
	for i := 0; i < len(list); i++ {
		template = append(template, list[i].(string))
	}
	return template
}

func (c Config) Tag() TagList {
	var tags TagList
	for _, tmp := range c["Tags"].([]interface{}) {
		for key, val := range tmp.(map[interface{}]interface{}) {
			k := key.(string)
			v := val.(string)
			t := &cf.Tag{Key: &k, Value: &v}
			tags = append(tags, t)
		}
	}
	return tags
}
