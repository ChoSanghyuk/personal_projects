package main

import (
	"fmt"
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func main() {
	nanjicamping()
}

func sample() {
	ok := hook.AddEvents("q", "ctrl", "shift")
	if ok {
		fmt.Println("add events...")
	}

	keve := hook.AddEvent("k")
	if keve {
		fmt.Println("you press... ", "k")
	}

	mleft := hook.AddEvent("mleft")
	if mleft {
		fmt.Println("you press... ", "mouse left button")
	}
}

/*
좌표는 마우스 트레이서의 논리 좌표 사용
한칸 내려진 이후의 상대적 상태에서 찾아야함

1. 팝업 닫기
2. 한칸 내리기
3. 요일 클릭
4. 예약하기 클릭
5. 인원
6. 신청자 정보와 동일
7. 이메일
8. 자동방지 클릭
9. 전체 동의
10. 예약하기


*/

func nanjicamping() {
	ok := hook.AddEvents(r.F1, "ctrl")
	if ok {
		fmt.Println("Phase 1")
		r.MoveClick(849, 239)  // 팝업 닫기
		r.ScrollDir(1, "down") // 한칸 내리기
		moveClick(519, 666)    // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)    // 예약하기 버튼
	}

	ok = hook.AddEvents(r.F2, "ctrl")
	if ok {
		fmt.Println("Phase 2")
		moveClick(1029, 196)   // 인원
		moveClick(935, 697)    // 신청자 정보와 동일
		r.ScrollDir(1, "down") // 한칸 내리기
		moveClick(485, 849)    // 로봇이 아닙니다
	}

	ok = hook.AddEvents(r.F3, "ctrl")
	if ok {
		fmt.Println("Phase 3")
		r.KeyTap("enter") // 인증되었습니다
		time.Sleep(200 * time.Millisecond)
		r.ScrollDir(10, "down") // 한칸 내리기
		moveClick(631, 691)     // 전체 동의
		// moveClick(1303, 423)    // 예약하기
	}

}

func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(300 * time.Millisecond)
	r.MoveClick(x, y, args...)
}
