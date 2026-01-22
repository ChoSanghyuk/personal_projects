package macro

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
	"time"

	r "github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
	hook "github.com/robotn/gohook"
)

func TestSample1(t *testing.T) {
	sample()
}

func TestSample2(t *testing.T) {
	sample2()
}

func TestWholeProcess(t *testing.T) {
	time.Sleep(3 * time.Second)
	CampitMac()
}

/******************************************* Individual function Test *************************************************/

func TestLocations(t *testing.T) {
	done := make(chan struct{})
	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
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

func TestMoveClick(t *testing.T) {
	time.Sleep(3 * time.Second)
	// r.KeyPress("f5")
	r.KeyTap("r", "cmd")
	// moveClick(1047, 776)
}

func TestKeyTap(t *testing.T) {

	time.Sleep(2 * time.Second)
	r.KeyPress(r.KeyP, r.Lctrl) // 인증되었습니다
}

func TestScrollDown(t *testing.T) {

	time.Sleep(2 * time.Second)
	r.ScrollDir(1800, "down") // 한칸 내리기
}

func TestGrabScrollRight(t *testing.T) {

	time.Sleep(2 * time.Second)
	x, y := r.Location()
	r.MouseDown()

	// 3. Move mouse to the right (e.g., +200 pixels) while holding the button
	r.Move(x-800, y) // optional: adjust speed

	// 4. Optional: Wait a bit if needed
	time.Sleep(30 * time.Millisecond)

	// 5. Release the mouse button
	r.MouseUp()
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

func TestServerTime(t *testing.T) {

	start := time.Now()
	time.Sleep(2 * time.Second)
	serverTime, err := getServerTime("https://waiting-site.yanolja.com/leisure/R53OE/")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	diff := serverTime.Sub(start)

	fmt.Printf("Server time: %v\n", serverTime)
	fmt.Printf("Local time:  %v\n", start)
	fmt.Printf("Diff:  %v\n", serverTime.Sub(start))
	if diff > 1*time.Second {
		println(int(diff.Seconds())) // 이정도 만큼 더 빨리 새로고침
	}
}

func TestTime(t *testing.T) {
	openTime, _ := time.ParseInLocation("2006-01-02 15:04", "2025-09-24 08:56", time.Local)
	now := time.Now()

	fmt.Printf("Now time: %v\n", now)
	fmt.Printf("Open time: %v\n", openTime)

	println(now.After(openTime))
}

/*
5,582
310,589
4,986
325,992
*/

func TestReenforceUntilTarget(t *testing.T) {
	time.Sleep(10 * time.Second)
	ReenforceUntilTarget(15)
}

func TestCaptureImage(t *testing.T) {
	time.Sleep(5 * time.Second)
	imgPath := "last_msg.png"
	r.SaveCapture(imgPath, 1, 430, 370, 560)
}

func TestCaptureText(t *testing.T) {
	client := gosseract.NewClient()
	defer client.Close()
	imgPath := "last_msg.png"

	// Set language to Korean
	client.SetLanguage("kor")

	// Set the image file path
	client.SetImage(imgPath)

	// Perform OCR
	text, err := client.Text()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 1. Define the regex pattern
	// This looks for "남은 골드:" followed by optional whitespace
	// and captures digits, commas, and dots.
	re := regexp.MustCompile(`남은 골드:\s*([\d,.]+)`)

	// 2. Find the match
	match := re.FindStringSubmatch(text)

	if len(match) > 1 {
		// match[0]	남은 골드: 1,250,000.50	The Full Match. It includes the label and the number.
		// match[1]	1,250,000.50	The First Capturing Group. This is only what was inside the () in your regex.
		rawNumber := match[1]

		// 3. Remove commas and periods
		// We use a Replacer for efficiency and readability
		replacer := strings.NewReplacer(",", "", ".", "")
		cleanNumber := replacer.Replace(rawNumber)

		fmt.Println("Extracted Number:", cleanNumber)

	}
	if strings.Contains(text, "+1") {
		fmt.Println("HELLO")
	}

	fmt.Println("Extracted Text:")
	fmt.Println(text)
}

func TestTextFromDrag(t *testing.T) {
	done := make(chan struct{})
	hook.Register(hook.KeyDown, []string{"enter"}, func(e hook.Event) {
		text := textFromDrag()
		fmt.Println(text)
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
