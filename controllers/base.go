package controllers

import (
	"github.com/astaxie/beego"
	"net"
	"net/http"
	"strings"
)

type NestPrepare interface {
	NestPrepare()
}

type baseController struct {
	beego.Controller
	ip string
}

func (b *baseController) Prepare() {
	b.ip = getIp(b.Ctx.Request)
	if app, ok := b.AppController.(NestPrepare); ok {
		app.NestPrepare()
	}
}

// 获取ip
func getIp(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
