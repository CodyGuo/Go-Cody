package main

import (
	"github.com/CodyGuo/logs"
	"github.com/go-vgo/robotgo"
)

func main() {
	defer robotgo.StopEvent()
	a := robotgo.AddEvent("a")
	if a == 0 {
		logs.Notice("press a ...")
	}
}
