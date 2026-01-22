package macro

import (
	"fmt"
	"strings"
	"time"

	r "github.com/go-vgo/robotgo"
	"github.com/otiai10/gosseract/v2"
)

func ReenforceUntilTarget(targetLevel int) {
	client := gosseract.NewClient()
	// Set language to Korean
	client.SetLanguage("kor")
	defer client.Close()
	imgPath := "last_msg.png"
	var lastgold string = ""
	noMoney := false

	typeMsg("/강화")
	for true {
		time.Sleep(4 * time.Second)
		r.SaveCapture(imgPath, 5, 350, 300, 550)

		// Set the image file path
		client.SetImage(imgPath)

		// Perform OCR
		text, err := client.Text()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		// fmt.Println(text)
		nowgold := ExtractTheRemainGold(text)
		fmt.Println(nowgold)

		// fmt.Println(text)
		if nowgold == "" {
			fmt.Println(text)
			continue
		} else if nowgold == lastgold {
			continue
		} else if strings.HasPrefix(nowgold, "9") { // 맨 끝 G가 6으로 인식
			noMoney = true
			break
		} else {
			lastgold = nowgold
		}

		if strings.Contains(text, fmt.Sprintf("+%d", targetLevel)) {
			fmt.Println("HELLO")
			break
		}

		typeMsg("/강화")
	}

	moveClick(45, 1018) // 메신저 칸
	if noMoney {
		r.TypeStr("no money!!!")
	} else {
		r.TypeStr(fmt.Sprintf("%d강 강화 완료!!!", targetLevel))
	}
	time.Sleep(1 * time.Second)
	r.KeyTap(r.Enter)
	r.KeyTap(r.Enter)

}

// 1. Define the regex pattern
// This looks for "남은 골드:" followed by optional whitespace
// and captures digits, commas, and dots.
func ExtractTheRemainGold(text string) string {
	// 2. Find the match
	match := re.FindStringSubmatch(text)

	if len(match) > 1 {
		// match[0]	남은 골드: 1,250,000.50	The Full Match. It includes the label and the number.
		// match[1]	1,250,000.50	The First Capturing Group. This is only what was inside the () in your regex.
		rawNumber := match[1]

		// 3. Remove commas and periods
		// We use a Replacer for efficiency and readability
		replacer := strings.NewReplacer(",", "", ".", "", " ", "", ":", "")
		cleanNumber := replacer.Replace(rawNumber)

		return cleanNumber

	}
	return ""
}
