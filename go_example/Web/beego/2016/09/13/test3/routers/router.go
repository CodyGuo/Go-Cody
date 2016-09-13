package routers

import (
	"github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/13/test3/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
