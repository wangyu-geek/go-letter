// 优先初始化与框架无关的配置
package setting

import (
	"leeter/pkg/initing"
	"path/filepath"
	"runtime"
)

var appPath string

func init() {
	mustFirstInit()
}

func Init() {
	otherStepInit()
}

// mustFirstInit 初始化与框架无关的配置
func mustFirstInit() {
	setAppPath()
}

// otherStepInit 初始化与框架相关配置
func otherStepInit() {
	initing.Init()
}

func setAppPath() {
	_, file, _, _ := runtime.Caller(0)
	appPath, _ = filepath.Abs(filepath.Dir(filepath.Join(file, "../.."+string(filepath.Separator))))
}

func getAppPath() string {
	return appPath
}

func GetAppPath() string {
	app := getAppPath()
	return app
}
