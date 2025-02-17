package main

import (
	"time"

	r "github.com/go-vgo/robotgo"
)

func main() {

	// for {
	// 	if time.Now().Minute() == 0 {
	// 		r.KeyPress("f5")
	// 		break
	// 	}
	// 	// time.Sleep(time.Mi)
	// }
	kintexcamping()
}

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(350 * time.Millisecond)
	r.MoveClick(x, y, args...)
}
