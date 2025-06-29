package macro

import (
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func Earth() {

	done := make(chan struct{})

	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		moveClick(1202, 193) // 25.05.01
		moveClick(1146, 422) // 25.05.01
		time.Sleep(waitTime)
		moveClick(1064, 625) // 25.05.01
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		moveClick(524, 435) // 25.05.01
		r.TypeStr("조상혁")
		moveClick(504, 474) // 25.05.01
		r.TypeStr("010-6289-2458")
		moveClick(488, 535) // 25.05.01
		moveClick(571, 616) // 25.05.01
		moveClick(432, 645) // 25.05.01
		moveClick(368, 711) // 25.05.01
	})

	s := hook.Start()
	go func() {
		<-hook.Process(s) // Process events
	}()

	<-done
}
