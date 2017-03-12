package main

import (
	"fmt"
	"os"

	"github.com/astaxie/beego/logs"
)

func main() {
	log := logs.NewLogger(1000)
	log.SetLogger("console", "")
	fmt.Print("hello world.\n")
	log.Trace("trace")
	os.Exit(0)
}
