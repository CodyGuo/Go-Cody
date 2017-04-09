package main

import (
	_ "github.com/CodyGuo/Go-Cody/go_example/Web/networkDiagnosis/routers"
	"github.com/astaxie/beego"
	_ "github.com/codyguo/godaemon"
)

func main() {
	beego.Run()
}
