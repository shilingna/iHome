package routers

import (
	"github.com/astaxie/beego"
	"iHome/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
