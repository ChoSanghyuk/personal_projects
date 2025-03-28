package macro

import (
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func CampitMac() {

	done := make(chan struct{})
	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		selectDateCamfit()

	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		pressAndMoveRightCampfit()
	})

	hook.Register(hook.KeyDown, []string{"3"}, func(e hook.Event) {
		moveClick(987, 1081) // 사이트 예약하기
	})

	hook.Register(hook.KeyDown, []string{"esc"}, func(e hook.Event) {
		close(done)
	})

	s := hook.Start()
	go func() {
		<-hook.Process(s) // Process events
	}()

	<-done

}

func selectDateCamfit() {
	r.ScrollDir(1800, "down") // 한칸 내리기
	moveClick(1027, 290)
	r.Move(904, 685)
	scrollDown(500)
	moveClick(904, 685) // 25.05.01
	moveClick(627, 758) // 25.05.04
	moveClick(836, 1077)
	moveClick(836, 1077)
}

func pressAndMoveRightCampfit() {
	x, y := 1045, 814
	r.Move(x, y)
	time.Sleep(30 * time.Millisecond)

	r.MouseDown()

	r.Move(x-800, y)
	time.Sleep(30 * time.Millisecond)

	r.MouseUp()
	moveClick(1014, 838)
}
