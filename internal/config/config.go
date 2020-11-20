package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

//DbConfig ....
type DbConfig struct {
	Driver string `yaml:"driver"`
}

//Config ...
type Config struct {
	DB      DbConfig `yaml:"db"`
	Version string   `yaml:"version"`
}

//LoadConfig ...
func LoadConfing(filename string) (*Config, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var c = &Config{}
	err = yaml.Umarshal(file, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
