package main

import (
	"flag"
	"fmt"
	"os/exec"
	"syscall"
	_ "uac"
)

var sw bool

func init() {
	flag.BoolVar(&sw, "on", true, "enable")
}

func main() {
	flag.Parse()
	err := disable(sw)
	fmt.Println(err)
}

func disable(ok bool) error {
	var on string = "disable"
	if ok {
		on = "enable"
	}
	cmd := exec.Command("cmd.exe", "/c", `netsh interface set interface "以太网" `+on)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	return cmd.Run()
}
