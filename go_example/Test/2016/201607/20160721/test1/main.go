package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "-t", "www.hupu.net")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
