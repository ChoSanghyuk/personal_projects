package config

import (
	_ "embed"
	"fmt"

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

	Email struct {
		SMTP   SMTP   `yaml:"smtp"`
		Target string `yaml:"target"`
	} `yaml:"email"`
}

type apiConfig struct {
	Url    string `yaml:"url"`
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
}

func (s SMTP) Server() string {
	return s.ServerI
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
