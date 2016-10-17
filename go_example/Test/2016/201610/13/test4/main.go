package main

import (
	// "fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "ping 127.0.0.1 -t")
	cmd.Stdout = os.Stdout
	cmd.Start()
	cmd.Wait()
}
