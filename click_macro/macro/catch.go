package macro

import r "github.com/go-vgo/robotgo"

func CatchMacroMac(isReal bool) {

	moveClick(1047, 732) // 날짜 선택.
	moveClick(951, 1010) // 날짜 선택

	if isReal {
		moveClick(983, 1077)
	} else {
		r.Move(983, 1077) // test 할때는 마우스 잘 이동되는지만. 본 예약때만 moveClick 사용
	}

}
