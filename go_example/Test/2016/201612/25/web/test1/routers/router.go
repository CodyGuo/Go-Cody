package routers

import (
	"github.com/CodyGuo/Go-Cody/go_example/Test/2016/201612/25/web/test1/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
}
