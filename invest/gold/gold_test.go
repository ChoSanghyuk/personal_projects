package gold

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {

	CallGoldApi()
}

func TestCrawl(t *testing.T) {

	rtn := Crawl()
	fmt.Println(rtn)
	assert.NotEmpty(t, rtn, "Crawl() should return a value")

}
