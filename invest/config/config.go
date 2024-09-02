package config

import (
	_ "embed"
	"invest/util"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configByte []byte

type Config struct {
	Api      map[string]apiConfig   `yaml:"api"`
	Crawl    map[string]crawlConfig `yaml:"crawl"`
	Telegram struct {
		ChatId string `yaml:"chatId"`
		Token  string `yaml:"token"`
	} `yaml:"telegram"`
}

type apiConfig struct {
	Url    string            `yaml:"url"`
	Header map[string]string `yaml:"header"`
}

type crawlConfig struct {
	Url     string `yaml:"url"`
	CssPath string `yaml:"css-path"`
}

func NewConfig() (*Config, error) {

	var ConfigInfo Config = Config{}

	err := yaml.Unmarshal(configByte, &ConfigInfo)
	if err != nil {
		return nil, err
	}

	// util.Decode(&ConfigInfo.Gold.API.ApiKey)
	util.Decode(&ConfigInfo.Telegram.ChatId)
	util.Decode(&ConfigInfo.Telegram.Token)

	return &ConfigInfo, nil
}

func (c Config) ApiInfo(target string) (url string, header map[string]string) {
	return c.Api[target].Url, c.Api[target].Header
}

func (c Config) CrawlInfo(target string) (url string, cssPath string) {
	return c.Crawl[target].Url, c.Crawl[target].CssPath
}
