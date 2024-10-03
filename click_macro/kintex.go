package main

import (
	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

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
