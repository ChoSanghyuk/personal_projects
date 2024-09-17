package main

import (
	"testing"
	"time"

	r "github.com/go-vgo/robotgo"
)

func TestScrollDown(t *testing.T) {
	time.Sleep(2 * time.Second)
	r.ScrollDir(8, "down") // 한칸 내리기
}
