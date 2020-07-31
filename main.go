package main

import (
	"github.com/astaxie/beego"
	_ "wangqingshui/routers"
	_ "github.com/astaxie/beego/session/redis"
)

func main() {
	beego.Run()
}

