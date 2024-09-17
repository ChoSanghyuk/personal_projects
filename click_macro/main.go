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

func kintexcamping() {

	nth := 1
	isSeq := false

	weekendXy := [][2]int{
		{224, 231},
		{224, 231},
		{224, 242},
		{224, 263},
		{224, 282},
	}

	ok := hook.AddEvent("0")
	if ok {
		moveClick(1375, 363) // 다음달 이동
		moveClick(1319, 889) // 예매하기
	}

	ok = hook.AddEvent("a") // 자동 입력 방지 입력 준비하기 위해 a 입력
	if ok {
		moveClick(513, 504) // 자동입력 방지 문자
	}

	ok = hook.AddEvent("enter")
	if ok {
		moveClick(674, 292) // 닫기

		// moveClick(169, 260) // 테스트용 169,260
		// _ = weekendXy[nth][0]
		moveClick(weekendXy[nth][0], weekendXy[nth][1]) // 요일 선택
		moveClick(220, 383)                             // 이용기간
		if !isSeq {
			moveClick(192, 433) //1박 2일
		} else {
			moveClick(183, 445) //2박 3일
		}

		moveClick(400, 550) // 아래 오토캠핑 구역 선택
		moveClick(631, 457) // 아래 오토캠핑 구역 11번
		moveClick(786, 621) // 다음 단계
	}

	// 예상처럼 두번째칸이 주말_고양시민_신분증지참일 경우
	ok = hook.AddEvent("0")
	if ok {
		moveClick(58, 274)  // 유형 선택
		moveClick(845, 619) // 다음 단계
		moveClick(136, 197) // 생년월일 칸
		r.TypeStr("940901")
		moveClick(485, 377) // 카드 선택
		moveClick(485, 565) // 신한카드
		moveClick(845, 619) // 다음 단계

	}

	ok = hook.AddEvent("0") // 마지막 단계
	if ok {
		moveClick(56, 550)  // 동의하기
		moveClick(770, 579) // 결제하기
	}
}

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(300 * time.Millisecond)
	r.MoveClick(x, y, args...)
}
