package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path"
	"testing"
	"time"

	"github.com/go-vgo/robotgo"
	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	hook "github.com/robotn/gohook"
	"github.com/vcaesar/imgo"
)

func TestRefresh(t *testing.T) {
	time.Sleep(3 * time.Second)
	robotgo.KeyTap("r", "cmd")
}

func TestLocation(t *testing.T) {

	hook.Register(hook.MouseDown, nil, func(e hook.Event) {
		x, y := robotgo.Location()
		fmt.Printf("pos: %d, %d\n", x, y)
	})

	hook.Register(hook.KeyDown, []string{robotgo.Esc}, func(e hook.Event) {
		hook.End()
	})

	s := hook.Start()
	<-hook.Process(s)
}

/*
시작점 : (230, 150) , (990, 150)
w : 630
h : 850
*/

func TestCapture(t *testing.T) {
	time.Sleep(2 * time.Second)

	spots := [][]int{{230, 180}, {990, 180}}
	w := 520
	h := 820

	for i := range spots {
		bit := robotgo.CaptureScreen(spots[i][0], spots[i][1], w, h)
		defer robotgo.FreeBitmap(bit)

		img := robotgo.ToImage(bit)
		imgo.Save(fmt.Sprintf("test%d.png", i+1), img)
	}

}

func TestAllFileNames(t *testing.T) {

	dir := "./book"
	root := os.DirFS(dir)

	pngFiles, err := fs.Glob(root, "*.png")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range pngFiles {
		fmt.Println(path.Join(dir, f))
	}
}

func TestPDF(t *testing.T) {
	imp, _ := pdfcpu.Import("form:A3, pos:c, s:1.0", types.POINTS)
	pdfcpu.ImportImagesFile([]string{"./book/img001.png", "./book/img002.png", "./book/img003.png"}, "out.pdf", imp, nil)
}

func TestMergeToPDH(t *testing.T) {
	mergeToPDH("./book", "이펙티브_러스트")
}
