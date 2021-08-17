// 测试脚本的初始化

package testting

import (
	"github.com/beego/beego/v2/server/web"
	"log"
	"leeter/pkg/setting"
)

func init() {
	appPath := setting.GetAppPath()
	log.Println("apppath is", appPath)
	web.TestBeegoInit(appPath)
	setting.Init()
}
