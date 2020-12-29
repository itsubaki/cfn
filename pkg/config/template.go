package config

import (
	"io/ioutil"
	"strings"

	"github.com/aws/aws-sdk-go/service/cloudformation"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Resources []Resource `yaml:"resources"`
	Tags      []Tag      `yaml:"tags"`
}

type Resource struct {
	Name       string     `yaml:"name"`
	Region     string     `yaml:"region"`
	Type       string     `yaml:"type"`
	Properties []Property `yaml:"properties"`
}

type Tag struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Property struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func Read(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return &Config{}, err
	}

	var config Config
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return &Config{}, err
	}

	return &config, nil
}

func StackName(arg, name string) string {
	tmp := strings.Replace(name, "/", "-", -1)
	suffix := strings.Replace(tmp, ".yml", "", -1)
	return arg + "-" + suffix
}

func (c *Config) Tag() []*cloudformation.Tag {
	var tags []*cloudformation.Tag
	for _, tag := range c.Tags {
		k := tag.Name
		v := tag.Value
		tags = append(tags, &cloudformation.Tag{Key: &k, Value: &v})
	}
	return tags
}

func (c *Config) Reverse() []Resource {
	var reverse []Resource
	list := c.Resources
	for i := 0; i < len(list); i++ {
		reverse = append(reverse, list[len(list)-1-i])
	}
	return reverse
}

func (t *Resource) Body() (string, error) {
	buf, err := ioutil.ReadFile(t.Type)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

func (t *Resource) Parameter() []*cloudformation.Parameter {
	var parameters []*cloudformation.Parameter
	for _, p := range t.Properties {
		k := p.Name
		v := p.Value
		parameters = append(parameters, &cloudformation.Parameter{ParameterKey: &k, ParameterValue: &v})
	}
	return parameters
}
