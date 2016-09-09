package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	// c.Data["Website"] = "beego.me"
	// c.Data["Email"] = "astaxie@gmail.com"
	// c.TplName = "index.tpl"
	c.Ctx.WriteString("appname: " + beego.AppConfig.String("appname") +
		"\nhttpport: " + beego.AppConfig.String("httpport") +
		"\nrunmode: " + beego.AppConfig.String("runmode"))

	beego.Trace("trace 1")
	beego.Info("info 1")
	beego.SetLevel(beego.LevelInformational)
	beego.Trace("trace 2")
	beego.Info("info 2")
}
