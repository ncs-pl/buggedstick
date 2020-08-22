package main

import (
	"github.com/go-yaml/yaml"
	"io/ioutil"
)

// Config is the required properties to Authenticate
type Config struct {
	Host     string `yaml:"host"`     // The SMTP host address
	Port     int    `yaml:"port"`     // The SMTP host port
	Password string `yaml:"password"` // The Sender password
	Email    string `yaml:"email"`    // The Sender email address
}

// GetConfig returns the Configuration file "buggedstick.config.yml"
func GetConfig() (*Config, error) {
	file, err := ioutil.ReadFile("./buggedstick.config.yml")

	if err != nil {
		return nil, err
	}

	conf := &Config{}

	if err := yaml.Unmarshal(file, &conf); err != nil {
		return nil, err
	}

	return conf, nil
}
