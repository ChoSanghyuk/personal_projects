package api

import (
	"invest/config"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGold(t *testing.T) {

	url := config.ConfigInfo.Gold.API.Url
	cssPath := config.ConfigInfo.Gold.API.ApiKey

	rtn, err := CallApi(url, cssPath)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, rtn)

	t.Log(rtn)
}
