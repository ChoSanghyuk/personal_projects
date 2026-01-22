package macro

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	r "github.com/go-vgo/robotgo"
)

var isMac bool = true
var t []int
var d []int
var macType []int = []int{45, 1018}
var macDrag []int = []int{350, 900, 5, 402}
var winType []int = []int{}
var winDrag []int = []int{}

var re *regexp.Regexp = regexp.MustCompile(`남은 골드.\s*([\d,.]+)`) // :\s => .

func init() {
	if isMac {
		t = macType
		d = macDrag
	} else {
		t = winType
		d = winDrag
	}
}

// ReenforceUntilTargetDrag drags text, pastes to terminal, and checks for target word
func ReenforceUntilTargetDrag(target int) {
	var old string = ""
	for {
		time.Sleep(3 * time.Second)
		text := textFromDrag()
		// Step 4: Check if target word exists in copied text

		if notYetCompleted(old, text) {
			continue
		} else {
			old = text
		}

		if !isHidden(text) {
			typeMsg("/판매")
			continue
		}

		if strings.Contains(text, fmt.Sprintf("→ +%d", target)) {
			fmt.Printf("Target word '%s' found! Stopping.\n", target)
			typeMsg(fmt.Sprintf("%d강 강화 완료", target))
			break
		}

		if noMoney(text, 10000000) {
			typeMsg("no money")
			break
		}

		// fmt.Printf("Target word '%s' not found. Continuing...\n", targetWord)
		typeMsg("/강화")

	}
}

func noMoney(text string, limit int) bool {
	match := re.FindStringSubmatch(text) // oldMatched
	if len(match) > 1 {
		rawNumber := match[1]
		replacer := strings.NewReplacer(",", "", ".", "", " ", "", ":", "")
		cleanNumber := replacer.Replace(rawNumber)
		value, err := strconv.Atoi(cleanNumber)
		if err != nil {
			fmt.Sprintln("잔여 gold 숫자 변환 오류", err.Error())
			return false
		}
		if value < limit {
			return true
		}
	}
	return false
}

func notYetCompleted(old string, new string) bool {

	om := re.FindStringSubmatch(old) // oldMatched
	nm := re.FindStringSubmatch(new) // newMatched

	if len(om) > 1 && len(nm) > 1 {
		return om[1] == nm[1]
	}
	return false
}

func isHidden(text string) bool {
	if strings.Contains(text, "→ +1") { // check
		if strings.Contains(text, "검") || strings.Contains(text, "몽둥이") || strings.Contains(text, "막대") {
			if strings.Contains(text, "광선검") {
				return true
			}
		} else {
			return true
		}
	}
	return false
}

func textFromDrag() string {

	// Step 1: Drag to select text in the specified region
	fmt.Println("Dragging text...")
	dragText(d[0], d[1], d[2], d[3]) // todo. fix
	time.Sleep(500 * time.Millisecond)

	// Copy the selected text
	r.KeyTap("c", "cmd") // Copy (macOS)
	time.Sleep(300 * time.Millisecond)

	// Step 2: Get the exact text from clipboard
	copiedText, err := r.ReadAll()
	if err != nil {

		return ""
	}
	fmt.Println("Copied text:", copiedText)

	// Step 3: Paste to terminal
	// moveClick(45, 1018) // Terminal position - adjust as needed
	// time.Sleep(300 * time.Millisecond)
	// r.KeyTap("v", "cmd") // Paste
	// r.KeyTap(r.Enter)
	// time.Sleep(1 * time.Second)

	return copiedText

}

// dragText performs a mouse drag from (x1, y1) to (x2, y2) to select text
func dragText(x1, y1, x2, y2 int) {
	r.Move(x1, y1)
	time.Sleep(200 * time.Millisecond)
	r.MouseDown("left")
	time.Sleep(200 * time.Millisecond)
	r.MoveSmooth(x2, y2)
	time.Sleep(200 * time.Millisecond)
	r.MouseUp("left")
}

func typeMsg(text string) {
	moveClick(t[0], t[1]) // 메신저 칸
	r.TypeStr(text)
	time.Sleep(1 * time.Second)
	r.KeyTap(r.Enter)
	r.KeyTap(r.Enter)
}
