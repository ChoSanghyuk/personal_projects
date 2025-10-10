package macro

import (
	"fmt"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func HwaDamMac() {

	done := make(chan struct{})
	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		r.ScrollDir(1700, "down")
		// moveClick(839,525) // 본 방
		moveClick(867, 1097) // 테스트
	})
	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		r.MoveClick(860, 1080) // 옵션 선택하기
	})

	hook.Register(hook.KeyDown, []string{"3"}, func(e hook.Event) {
		r.MoveClick(941, 550) // 날짜 선택. 4째주 금요일
	})

	hook.Register(hook.KeyDown, []string{"4"}, func(e hook.Event) {
		r.MoveClick(1005, 601) // 시간 선택. 11:40
	})

	hook.Register(hook.KeyDown, []string{"5"}, func(e hook.Event) {
		// r.MoveClick(844, 583) // 입장권 성인
		r.MoveClick(869, 440) // 입장권 성인
	})

	hook.Register(hook.KeyDown, []string{"6"}, func(e hook.Event) {
		r.MoveClick(807, 675) // 일반
	})

	hook.Register(hook.KeyDown, []string{"7"}, func(e hook.Event) {
		for i := 1; i < 5; i++ {
			r.MoveClick(702, 764) // 수량
		}
	})
	hook.Register(hook.KeyDown, []string{"8"}, func(e hook.Event) {
		r.MoveClick(980, 964) // 모노레일 선택
	})

	hook.Register(hook.KeyDown, []string{"9"}, func(e hook.Event) {
		r.KeyTap("enter") // enter
	})

	/* phase 2 모노레일 */
	hook.Register(hook.KeyDown, []string{"q"}, func(e hook.Event) {
		r.MoveClick(883, 491) // 1구간 선택
	})
	hook.Register(hook.KeyDown, []string{"w"}, func(e hook.Event) {
		r.MoveClick(693, 518) // 이용 시간 12:00
	})
	hook.Register(hook.KeyDown, []string{"e"}, func(e hook.Event) {
		r.MoveClick(693, 579) // 모노레일 성인
	})
	hook.Register(hook.KeyDown, []string{"r"}, func(e hook.Event) {
		for i := 1; i < 5; i++ {
			r.MoveClick(702, 741) // 수량
		}
	})
	hook.Register(hook.KeyDown, []string{"t"}, func(e hook.Event) {
		r.MoveClick(993, 970) // 모노레일 성인
	})

	hook.Register(hook.KeyDown, []string{"."}, func(e hook.Event) {
		x, y := r.Location()
		fmt.Printf("%d,%d\n", x, y)
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
