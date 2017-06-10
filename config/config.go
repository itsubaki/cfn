package config

import (
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Config map[interface{}]interface{}

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
