package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type dotaConfig struct {
	DotaAPIPrefix string `yaml:"api_prefix"`
	DotaAPIKey    string `yaml:"api_key"`
}

func loadDotaConfig(filePath string) (*dotaConfig, error) {
	bytes, err := ioutil.ReadFile(filePath + "/opendotaapi.yaml")
	if err != nil {
		return nil, err
	}
	var config dotaConfig
	err = yaml.Unmarshal(bytes, &config)
	return &config, err
}
