package main

import (
	"macro/macro"
	"time"

	r "github.com/go-vgo/robotgo"
)

func main() {

	for {
		if time.Now().Minute() == 0 {
			r.KeyPress("f5")
			break
		}
		// time.Sleep(time.Millisecond)
	}
	macro.Earth()
}
