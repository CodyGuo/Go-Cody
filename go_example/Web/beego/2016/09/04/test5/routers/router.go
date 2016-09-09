package routers

import (
	"github.com/CodyGuo/Go-Cody/go_example/Web/beego/2016/09/04/test5/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
