package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/go-vgo/robotgo"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/imgo"
)

func main() {

	// 기본 경로
	path := "./book"

	// 책 설정 값
	start := 1
	end := 50 // 245
	name := "이펙티브_러스트"

	// 캡처 설정 값. TestCapture에서 값 확인
	spots := [][2]int{{230, 180}, {990, 180}}
	w := 520
	h := 820

	// 코드 시작
	ok := hook.AddEvent(robotgo.Enter)
	if ok {
		fmt.Println("Event Start")
	}

	for i := start; i <= end; i += 2 {
		capture(spots[0], w, h, fmt.Sprintf("%s/img%03d.png", path, i))
		capture(spots[1], w, h, fmt.Sprintf("%s/img%03d.png", path, i+1))
		move()
	}

	mergeToPDH(path, name)

}

func capture(spot [2]int, w, h int, name string) {

	bit := robotgo.CaptureScreen(spot[0], spot[1], w, h)
	defer robotgo.FreeBitmap(bit)

	img := robotgo.ToImage(bit)
	imgo.Save(name, img)
}

func move() {
	robotgo.MoveClick(1613, 573)
}

func mergeToPDH(path string, name string) {

	root := os.DirFS(path)

	pngFiles, err := fs.Glob(root, "*.png")
	if err != nil {
		log.Fatal(err)
	}

	files := make([]string, len(pngFiles))
	for i, f := range pngFiles {
		files[i] = fmt.Sprintf("%s/%s", path, f)
	}

	imp, _ := pdfcpu.Import("form:A3, pos:c, s:1.0", types.POINTS)
	pdfcpu.ImportImagesFile(files, name+".pdf", imp, nil)
}
