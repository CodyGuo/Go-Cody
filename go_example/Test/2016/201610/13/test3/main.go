package main

import (
	// "fmt"
	"os/exec"
	"syscall"
)

func main() {
	cmd := exec.Command("cmd", "/c", "start http://127.0.0.1")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
	for {
	}
}
