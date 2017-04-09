package routers

import (
	"github.com/CodyGuo/Go-Cody/go_example/Web/networkDiagnosis/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/ws", &controllers.WsPing{})
	beego.Router("/", &controllers.MainController{})
}
