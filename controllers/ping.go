package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"leeter/serializer"
)

type PingController struct {
	web.Controller
}

func (c *PingController) Get() {
	pong := serializer.BuildPong("yell")
	resp := serializer.Response{
		Data:  pong,
	}
	c.SetData(resp)
	c.ServeJSON()
}