package main

import (
	"os/exec"
)

func main() {
	exec.Command("start", "http://www.baidu.com").Run()
}
