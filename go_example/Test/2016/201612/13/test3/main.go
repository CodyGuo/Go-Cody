package main

import (
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("/usr/bin/ssh cisco@10.10.2.252")
	// cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	cmd.Start()
	cmd.Wait()
}
