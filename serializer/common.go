package serializer

import (
	"github.com/beego/beego/v2/server/web"
)

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error,omitempty"`
}

const (
	//CodeParamErr 参数错误
	CodeParamErr = 40001
	//InternalErr 内部错误
	InternalErr = 50001
)

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	res := Response{
		Status: errCode,
		Msg:    msg,
	}
	// 生产环境隐藏底层报错
	if err != nil && web.BConfig.RunMode == web.DEV {
		res.Error = err.Error()
	}
	return res
}

// ParamErr 各种参数错误
func ParamErr(msg string, err error) Response {
	if msg == "" {
		msg = "参数错误"
	}
	return Err(CodeParamErr, msg, err)
}

// BuildValidateErr 校验参数错误
//func BuildValidateErr(validErrors []*validation.Error) Response {
//	for _, err := range validErrors {
//		log.Println(err.Key, err.Message)
//	}
//	var msg string
//	if len(validErrors) > 0 {
//		err := validErrors[0]
//		msg = err.Key + err.Message
//	}
//
//	return ParamErr(msg, errors.New(msg))
//}
