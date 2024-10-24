package main

import (
	"fmt"
	"testing"
	"time"

	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

func TestSample1(t *testing.T) {
	sample()
}

func TestSample2(t *testing.T) {
	sample2()
}

func TestSaveBlog(t *testing.T) {
	time.Sleep(3 * time.Second)
	saveBlogAsPdf()
}

/******************************************* Individual function Test *************************************************/

func TestKeyTap(t *testing.T) {

	time.Sleep(2 * time.Second)
	r.KeyPress(r.KeyP, r.Lctrl) // 인증되었습니다
}

func TestScrollDown(t *testing.T) {

	time.Sleep(2 * time.Second)
	r.ScrollDir(3, "down") // 한칸 내리기
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

// 특정 좌표의 color을 가져오는 기능은 더 연구가 필요함
func TestGetLocationColor(t *testing.T) {
	time.Sleep(3 * time.Second)

	fmt.Println(r.GetPixelColor(1884, 572))
	fmt.Println(r.GetPixelColor(1887, 658))
}

func TestGetMousePos(t *testing.T) {
	time.Sleep(3 * time.Second)
	x, y := r.Location()

	// Print the mouse coordinates
	fmt.Printf("Mouse is at (%d, %d)\n", x, y)
}

func TestRegisterEnter(t *testing.T) {
	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		fmt.Println("Hello")
	})
	s := hook.Start()
	<-hook.Process(s)

	time.Sleep(10 * time.Second)
}
