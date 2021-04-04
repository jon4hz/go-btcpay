package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BTCPay struct {
		URL      string `yaml:"url"`
		Username string `yaml:"username,omitempty"`
		Password string `yaml:"password,omitempty"`
		APIKey   string `yaml:"apiKey,omitempty"`
	} `yaml:"BTCPayServer"`
}

func ReadConf(filename string) (*Config, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %v", filename, err)
	}
	return c, nil
}
