package config

import (
	"testing"
)

func TestConfigNew(t *testing.T) {

	conf, _ := NewConfig()
	t.Log(*conf.Key.KIS["appkey"])
	t.Log(*conf.Key.KIS["appsecret"])

}
