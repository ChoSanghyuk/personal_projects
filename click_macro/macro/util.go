package macro

import (
	"time"

	r "github.com/go-vgo/robotgo"
)

var waitTime = 300 * time.Millisecond

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(waitTime)
	r.MoveClick(x, y, args...)
}

func scrollDown(d int) {
	time.Sleep(waitTime)
	r.ScrollDir(d, "down")
}
