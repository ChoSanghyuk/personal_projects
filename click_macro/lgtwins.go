package main

import (
	hook "github.com/robotn/gohook"
)

func lgtwinsCns() {
	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		moveClick(289, 529) // 요일 선택
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		moveClick(734, 413) // 수량
		moveClick(715, 464) // 2매
		moveClick(852, 723) // 예매
	})

	s := hook.Start()
	<-hook.Process(s)
}

func lgtwins() {
	x, y := 898, 560

	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		moveClick(x, y)     // 예매
		moveClick(625, 703) //예매안내 확인

	})

	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) { // todo. enter 맞나?
		moveClick(509, 661) // test. 279,662
		moveClick(457, 587) //
		moveClick(444, 440) // 자리1 선택
		moveClick(459, 436) // 자리2 선택
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		moveClick(915, 841) // 다음단계
	})

	s := hook.Start()
	<-hook.Process(s)
}

func lgtwinsTest() {

	x, y := 898, 460

	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		moveClick(x, y)     // 예매
		moveClick(625, 703) //예매안내 확인

	})

	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) { // todo. enter 맞나?
		moveClick(279, 662)
		moveClick(457, 587) //
		moveClick(302, 592) // 자리1 선택
		moveClick(317, 592) // 자리2 선택
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		moveClick(915, 841) // 다음단계
	})

	s := hook.Start()
	<-hook.Process(s)
}
