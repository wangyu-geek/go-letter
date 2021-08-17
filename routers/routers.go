package routers

import (
	"github.com/beego/beego/v2/server/web"
	"leeter/controllers"
)

func init() {
	web.Router("/ping", &controllers.PingController{})
	web.SetStaticPath("/statics", "statics")
}
