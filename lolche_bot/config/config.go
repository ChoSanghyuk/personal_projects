package config

import (
	_ "embed"

	"gopkg.in/yaml.v3"
)

//go:embed config.yaml
var configByte []byte

type Config struct {
	TeleBot struct {
		Token  string `yaml:"token"`
		ChatId string `yaml:"chatId"`
	} `yaml:"telegram"`
}

func NewConfig() (*Config, error) {
	var ConfigInfo Config = Config{}

	err := yaml.Unmarshal(configByte, &ConfigInfo)
	if err != nil {
		return nil, err
	}
	return &ConfigInfo, nil
}

func (c Config) Telebot() (token string, chatId string) {
	return c.TeleBot.Token, c.TeleBot.ChatId
}
