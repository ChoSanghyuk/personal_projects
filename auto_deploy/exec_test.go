package main

import (
	"log"
	"os/exec"
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := exec.Command("./deploy_invest_indicator.sh")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
