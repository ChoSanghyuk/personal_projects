package event

import (
	"fmt"
	"testing"
)

type temp struct {
	filedA string
}

func TestTemp(t *testing.T) {

	li := make([]temp, 10)
	a := &li[0]
	a.filedA = "변경 후"

	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v", li)
}
