package main

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (self *HomeController) Get() {
	self.Ctx.WriteString("hello world.")
}

func main() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
