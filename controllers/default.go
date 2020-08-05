package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"wangqingshui/library/monitor"
)

type MainController struct {
	beego.Controller
	monitor *monitor.Collect
}

func (c *MainController) Prepare() {
	controller, action := c.GetControllerAndAction()
	c.monitor = monitor.NewCollect(monitor.MODULE_REQUEST, controller+"."+action, nil)
}
func (c *MainController) Finish() {
	c.monitor.Push()
}

func (c *MainController) Get() {

	//c.SetSession("uid", 10)
	//c.Ctx.SetCookie("username", "junmo")
	fmt.Println(333)
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
	c.Render()
}
