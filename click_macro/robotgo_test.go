package main

import (
	"fmt"
	"testing"
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

/******************************************* Scenario Test *************************************************/
func TestKintex(t *testing.T) {
	kintexcamping()
}

/******************************************* Individual function Test *************************************************/
func TestScrollDown(t *testing.T) {

	time.Sleep(2 * time.Second)
	r.ScrollDir(1, "down") // 한칸 내리기
}

func TestRightClickEvent(t *testing.T) {
	ok := hook.AddMouse(r.Mright)
	if ok {
		fmt.Println("HI")
	}
}

func TestTypeStr(t *testing.T) {
	time.Sleep(2 * time.Second)
	moveClick(136, 197) // 생년월일 칸
	r.TypeStr("940901")
}

func TestTemp(t *testing.T) {
	ok := hook.AddEvent("0")
	if ok {
		moveClick(58, 274)  // 유형 선택
		moveClick(845, 619) // 다음 단계
		moveClick(136, 197) // 생년월일 칸
		r.TypeStr("940901")
		moveClick(485, 377) // 카드 선택
		moveClick(485, 565) // 신한카드
		moveClick(845, 619) // 다음 단계

	}
}
