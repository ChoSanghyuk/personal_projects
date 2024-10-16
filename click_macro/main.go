package main

import (
	"time"

	r "github.com/go-vgo/robotgo"
)

func main() {
	lgtwins()
}

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(350 * time.Millisecond)
	r.MoveClick(x, y, args...)
}
