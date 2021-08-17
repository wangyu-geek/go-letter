package main

import (
	"github.com/beego/beego/v2/server/web"
	"log"
	"leeter/pkg/setting"
	_ "leeter/pkg/setting" //初始化配置
	_ "leeter/routers"     //初始化路由
)

func main() {
	setting.Init()
	log.Println("app start")
	web.Run()
}