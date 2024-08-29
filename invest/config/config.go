package config

import (
	_ "embed"
	"fmt"
	"invest/util"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configByte []byte

type Config struct {
	Gold struct {
		API   apiConfig   `yaml:"api"`
		Crawl crawlConfig `yaml:"crawl"`
		Bound bound       `yaml:"bound"`
	} `yaml:"gold"`
	Bitcoin struct {
		API   apiConfig   `yaml:"api"`
		Crawl crawlConfig `yaml:"crawl"`
		Bound bound       `yaml:"bound"`
	} `yaml:"bitcoin"`
	RealEstate struct {
		Crawl crawlConfig `yaml:"crawl"`
	} `yaml:"estate"`

	Email struct {
		SMTP   SMTP   `yaml:"smtp"`
		Target string `yaml:"target"`
	} `yaml:"email"`

	Telegram struct {
		ChatId string `yaml:"chatId"`
		Token  string `yaml:"token"`
	} `yaml:"telegram"`
}

type apiConfig struct {
	Url    string `yaml:"url"`
	ID     string `yaml:"id"`
	ApiKey string `yaml:"api-key"`
}

type crawlConfig struct {
	Url     string `yaml:"url"`
	CssPath string `yaml:"css-path"`
}

type bound struct {
	Lower float64 `yaml:"lower"`
	Upper float64 `yaml:"upper"`
}

type SMTP struct {
	ServerI   string `yaml:"server"`
	PortI     string `yaml:"port"`
	UserI     string `yaml:"user"`
	PasswordI string `yaml:"password"`
}

var ConfigInfo Config = Config{}

func init() {

	err := yaml.Unmarshal(configByte, &ConfigInfo)
	if err != nil {
		fmt.Println(err)
	}

	util.Decode(&ConfigInfo.Email.SMTP.PasswordI)
	util.Decode(&ConfigInfo.Gold.API.ApiKey)
	util.Decode(&ConfigInfo.Bitcoin.API.ID)
	util.Decode(&ConfigInfo.Bitcoin.API.ApiKey)

	util.Decode(&ConfigInfo.Telegram.ChatId)
	util.Decode(&ConfigInfo.Telegram.Token)

}

func (s SMTP) Server() string {
	return s.ServerI
}

func NewConfigInfo() *Config {
	return &ConfigInfo
}

func (s SMTP) Port() string {
	return s.PortI
}

func (s SMTP) User() string {
	return s.UserI
}

func (s SMTP) Password() string {
	return s.PasswordI
}
