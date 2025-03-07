package main

import (
	"macro/macro"
	"time"
)

func main() {

	for {
		if time.Now().Minute() == 0 {
			// r.KeyPress("f5")
			break
		}
		// time.Sleep(time.Millisecond)
	}
	macro.CatchMacroMac(true)
}
