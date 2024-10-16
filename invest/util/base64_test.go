package util

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {

	e := base64.StdEncoding.EncodeToString([]byte("7312714018"))
	fmt.Println(e)

	d, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(d))
}
