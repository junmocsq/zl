package routers

import (
	"github.com/astaxie/beego"
	"wangqingshui/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
