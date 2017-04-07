package routers

import (
	"github.com/CodyGuo/Go-Cody/go_example/Test/2016/201612/25/beego1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ws", &controllers.WsPing{})
	beego.Router("/", &controllers.MainController{})
}
