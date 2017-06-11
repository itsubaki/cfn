package config

import (
	"io/ioutil"

	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	yaml "gopkg.in/yaml.v2"
)

type Config map[string]interface{}
type TemplateList []interface{}

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
	return c["Templates"].([]interface{})
}

func (c Config) Tag() []*cf.Tag {
	var tags []*cf.Tag
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
