package main

import (
	"fmt"
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

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

func nanjicampingSeq() {

	var date_x, date_y = 749, 608 // 10월 12일

	ok := hook.AddEvent("0")
	if ok {
		fmt.Println("Phase 1")
		r.MoveClick(849, 239)     // 팝업 닫기
		r.ScrollDir(1, "down")    // 한칸 내리기
		moveClick(date_x, date_y) // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)       // 예약하기 버튼
	}

	ok = hook.AddEvent("0")
	if ok {
		fmt.Println("Phase 2")
		moveClick(1029, 196)   // 인원
		moveClick(935, 697)    // 신청자 정보와 동일
		r.ScrollDir(1, "down") // 한칸 내리기
		moveClick(485, 849)    // 로봇이 아닙니다
	}

	ok = hook.AddEvent("0")
	if ok {
		fmt.Println("Phase 3")
		r.KeyTap("enter") // 인증되었습니다
		time.Sleep(200 * time.Millisecond)
		r.ScrollDir(10, "down") // 한칸 내리기
		moveClick(631, 691)     // 전체 동의
		r.Move(1207, 388)
	}

}

func nanjicamping() {

	// test 673, 675
	var date_x, date_y = 672, 673 // 10월 12일

	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		fmt.Println("Phase 1")
		r.MoveClick(849, 239)     // 팝업 닫기
		r.ScrollDir(1, "down")    // 한칸 내리기
		moveClick(date_x, date_y) // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)       // 예약하기 버튼
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		fmt.Println("Phase 2")
		moveClick(1029, 196)   // 인원
		moveClick(935, 685)    // 신청자 정보와 동일 935,345
		r.ScrollDir(1, "down") // 한칸 내리기
		moveClick(485, 849)    // 로봇이 아닙니다
	})

	hook.Register(hook.KeyDown, []string{"3"}, func(e hook.Event) {
		if e.Keychar == 13 { // 이벤트 직후 enter가 눌리면 중복 실행되는 오류 수정
			return
		}
		fmt.Println("Phase 3")
		r.KeyTap("enter") // 인증되었습니다
		time.Sleep(200 * time.Millisecond)
		r.ScrollDir(10, "down") // 한칸 내리기
		moveClick(631, 691)
		// 전체 동의
		r.Move(1207, 388)
	})

	hook.Register(hook.KeyDown, []string{"4"}, func(e hook.Event) { // sub task. 팝업 닫은 상태에서의 로직 수행
		fmt.Println("Phase 1")
		// r.MoveClick(849, 239)     // 팝업 닫기
		r.ScrollDir(1, "down")    // 한칸 내리기
		moveClick(date_x, date_y) // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)       // 예약하기 버튼
	})

	s := hook.Start()
	<-hook.Process(s)
}

func nanjibbq() {

	// test 673, 675
	var date_x, date_y = 747, 670 // 10월 12일

	hook.Register(hook.KeyDown, []string{"1"}, func(e hook.Event) {
		fmt.Println("Phase 1")
		r.MoveClick(849, 239)     // 팝업 닫기
		r.ScrollDir(1, "down")    // 한칸 내리기
		moveClick(date_x, date_y) // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)       // 예약하기 버튼
	})

	hook.Register(hook.KeyDown, []string{"2"}, func(e hook.Event) {
		fmt.Println("Phase 2")
		moveClick(867, 938) // 회차선택
		time.Sleep(200 * time.Millisecond)
		r.ScrollDir(3, "down") // 3칸 내리기
		moveClick(1029, 371)   // 인원
		moveClick(935, 865)    // 신청자 정보와 동일 935,345
		r.ScrollDir(1, "down") // 한칸 내리기
		moveClick(485, 1017)   // 로봇이 아닙니다
	})

	hook.Register(hook.KeyDown, []string{"3"}, func(e hook.Event) {
		if e.Keychar == 13 { // 이벤트 직후 enter가 눌리면 중복 실행되는 오류 수정
			return
		}
		fmt.Println("Phase 3")
		r.KeyTap("enter") // 인증되었습니다
		time.Sleep(200 * time.Millisecond)
		r.ScrollDir(10, "down") // 한칸 내리기
		moveClick(631, 691)
		// 전체 동의
		r.Move(1207, 388)
	})

	hook.Register(hook.KeyDown, []string{"4"}, func(e hook.Event) { // sub task. 팝업 닫은 상태에서의 로직 수행
		fmt.Println("Phase 1")
		// r.MoveClick(849, 239)     // 팝업 닫기
		r.ScrollDir(1, "down")    // 한칸 내리기
		moveClick(date_x, date_y) // 한칸 내린 상태에서 요일 클릭
		moveClick(945, 792)       // 예약하기 버튼
	})

	s := hook.Start()
	<-hook.Process(s)
}
