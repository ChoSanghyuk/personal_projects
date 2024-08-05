package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompare(t *testing.T) {

	url1 := "https://search.naver.com/search.naver?where=nexearch&sm=top_hty&fbm=0&ie=utf8&query=%ED%95%9C%EA%B5%AD%EA%B1%B0%EB%9E%98%EC%86%8C+%EC%8B%A4%EC%8B%9C%EA%B0%84+%EA%B8%88+%EC%8B%9C%EC%84%B8"
	url2 := ConfigInfo.Gold.Crawl.Url

	assert.Equal(t, url1, url2, "URLs should be equal")

}

func TestCompare2(t *testing.T) {

	path1 := "#main_pack > section.sc_new.pcs_gold_rate._cs_gold_rate > div > div.gold_price.up > a > strong"
	path2 := ConfigInfo.Gold.Crawl.CssPath

	assert.Equal(t, path1, path2, "path should be equal")

}

func TestDecoded(t *testing.T) {

	t.Log(ConfigInfo.Gold.API.ApiKey)

}
