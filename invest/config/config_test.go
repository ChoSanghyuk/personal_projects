package config

import (
	"testing"
)

func TestConfigNew(t *testing.T) {

	conf, _ := NewConfig()
	t.Log(conf.Crawl["estate"].Url)

}
