package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("taskkill", "/F", "/IM", "cmd.exe")
	cmd.Run()
}
