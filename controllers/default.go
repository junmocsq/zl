package controllers

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.SetSession("uid", 10)
	c.Ctx.SetCookie("username", "junmo")

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
