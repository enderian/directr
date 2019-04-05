package types

import (
	"github.com/codegangsta/cli"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadConfiguration(c *cli.Context) (*Configuration, error) {
	conf := &Configuration{}
	confBytes, err := ioutil.ReadFile(c.String("config"))
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(confBytes, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
