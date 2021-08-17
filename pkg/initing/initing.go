// 初始化系统配置
package initing

import (
	"leeter/pkg/aliyun/oss"
	"leeter/pkg/db"
	"leeter/pkg/logging"
)

// Init 初始化系统所需组件
func Init() {
	logging.Init() //初始化log
	oss.Init()     //初始化oss
	db.Init()      //初始化数据库
}
