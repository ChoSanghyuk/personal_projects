package main

import (
	"fmt"
	"time"

	r "github.com/go-vgo/robotgo"
)

func saveBlogAsPdf() {

	from := 3
	to := 65
	format := "https://dev-ote.tistory.com/%d"

	for i := from; i <= to; i++ {
		url := fmt.Sprintf(format, i)

		r.MoveClick(485, 65)
		r.TypeStr(url)
		r.KeyTap("enter") // 이동

		time.Sleep(2 * time.Second)
		r.KeyPress(r.KeyP, r.Lctrl)
		time.Sleep(1 * time.Second)
		r.KeyTap("enter") // 이동
		time.Sleep(1 * time.Second)
		r.MoveClick(779, 597)
		time.Sleep(500 * time.Millisecond)
		r.KeyPress(r.Home)
		time.Sleep(500 * time.Millisecond)
		r.TypeStr(fmt.Sprintf("%02d_", i))
		time.Sleep(500 * time.Millisecond)
		r.KeyTap("enter") // 이동
		time.Sleep(1 * time.Second)
	}

}
