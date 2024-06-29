package config

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configByte []byte

type Config struct {
	GoldConfig struct {
		API   apiConfig   `yaml:"api"`
		Crawl crawlConfig `yaml:"crawl"`
	} `yaml:"gold"`
}

type apiConfig struct {
	Url    string `yaml:"url"`
	ApiKey string `yaml:"api-key"`
}

type crawlConfig struct {
	Url     string `yaml:"url"`
	CssPath string `yaml:"css-path"`
}

var ConfigInfo Config = Config{}

func init() {

	err := yaml.Unmarshal(configByte, &ConfigInfo)
	if err != nil {
		fmt.Println(err)
	}
}
