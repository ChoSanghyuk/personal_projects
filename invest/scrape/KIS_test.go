package scrape

import (
	"invest/config"
	"testing"
)

func TestGenerateToken(t *testing.T) {

	conf, _ := config.NewConfig()
	token, err := GenerateToken(conf.KisAppKey(), conf.KisAppSecret())
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
}

func TestKISApi(t *testing.T) {

	conf, _ := config.NewConfig()
	KISApi(conf.KisAppKey(), conf.KisAppSecret())
}
