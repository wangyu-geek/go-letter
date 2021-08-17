package oss

import (
	"github.com/beego/beego/v2/server/web"
	"log"
	"leeter/pkg/logging"
)

var ossConfig Config

func Init() {
	initOSS()
}

func initOSS() {
	oss, err := web.AppConfig.GetSection("aliyun_oss")
	if err != nil {
		logging.ErrorConfig("aliyun oss init()", "aliyun", oss, "GetSection fail", err)
		return
	}
	ossConfig.EndPoint = oss["endpoint"]
	ossConfig.AccessKey = oss["accesskey"]
	ossConfig.AccessKeySecret = oss["accesskeysecret"]
	ossConfig.Domain = oss["domain"]
	ossConfig.Bucket = oss["bucket"]

	log.Printf("init oss config %v\n", ossConfig)
}
