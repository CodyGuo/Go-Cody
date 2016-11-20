package main

import (
	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(1000)
	log.SetLogger("console", "")
	log.Async(12)
	log.Trace("trace")
}
