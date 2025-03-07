package macro

import (
	"time"

	r "github.com/go-vgo/robotgo"
)

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(300 * time.Millisecond)
	r.MoveClick(x, y, args...)
}
