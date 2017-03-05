package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	// log := logs.NewLogger(1000)
	// log.SetLogger("console", "")
	//
	// log.Info("hello %s\n", "world.")
	// log.Debug("hi %s", "codguo")
	logs.Info("default %s\n", "info")
}
