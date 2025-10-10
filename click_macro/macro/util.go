package macro

import (
	"fmt"
	"net/http"
	"time"

	r "github.com/go-vgo/robotgo"
)

var waitTime = 300 * time.Millisecond

/****************************************************** Inner Function *************************************************/
func moveClick(x int, y int, args ...interface{}) {
	time.Sleep(waitTime)
	r.MoveClick(x, y, args...)
}

func scrollDown(d int) {
	time.Sleep(waitTime)
	r.ScrollDir(d, "down")
}

func getServerTime(url string) (time.Time, error) {
	resp, err := http.Head(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()

	dateHeader := resp.Header.Get("Date")
	if dateHeader == "" {
		return time.Time{}, fmt.Errorf("no Date header found")
	}

	// Parse the HTTP date format
	serverTime, err := time.Parse(time.RFC1123, dateHeader)
	if err != nil {
		return time.Time{}, err
	}

	return serverTime, nil
}
