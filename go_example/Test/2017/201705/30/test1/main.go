package main

import (
	"os"
	"os/exec"
	"strings"
)

func main() {
	name := "ping"
	args := "-n 10 www.qq.com"
	arg := strings.Split(args, " ")

	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}
