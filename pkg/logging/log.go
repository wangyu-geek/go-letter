package logging

import (
	"fmt"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const DefaultLogFormat = `{"filename":"%s/%s.%s","separate":["error", "warning", "notice", "info", "debug"]}`

type logSet struct {
	LogSavePath string
	LogSaveExt  string
	LogConfig   string
}

var LogSetting *logSet

func Init() {
	initLog()
	setLog()
}

func initLog() {
	LogSetting = &logSet{}
	LogSetting.LogSavePath, _ = web.AppConfig.String("LogSavePath")
	LogSetting.LogSaveExt, _ = web.AppConfig.String("LogSaveExt")
	appName := web.BConfig.AppName
	LogSetting.LogConfig = fmt.Sprintf(DefaultLogFormat, LogSetting.LogSavePath, appName, LogSetting.LogSaveExt)
}

var log *logs.BeeLogger

func setLog() {
	log = logs.NewLogger(10000)
	log.SetLogger(logs.AdapterMultiFile, LogSetting.LogConfig)
}

func Warn(format string, v ...interface{}) {
	log.Warn(format, v)
}

func Notice(format string, v ...interface{}) {
	log.Notice(format, v)
}

func Info(format string, v ...interface{}) {
	log.Info(format, v)
}

func Error(k string, method string, param interface{}, ret interface{}, msg string, err error) {
	log.Error(DefaultFormat, k, method, param, ret, msg, err)
}

func Param(method string, param interface{}, ret interface{}) {
	log.Warn(DefaultFormat, ParamError, method, param, ret, "", "")
}

func ErrorConfig(method string, param interface{}, ret interface{}, msg string, err error) {
	log.Error(ConfigError, method, param, ret, msg, err)
}
