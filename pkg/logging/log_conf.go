package logging

//定义基础错误关键字
const (
	DbError         = "mysql error"
	RedisError      = "redis error"
	RedisTokenError = "redisToken err"
	ServiceError    = "service error"
	ParamError      = "params error"
	ConfigError     = "config error"
	RmqError        = "rmq error"
	TokenError      = "token error"
	TraceError      = "trace error"
)

const (
	DefaultFormat = "k[%v] method[%v] param[%v] ret[%v] msg[%v] err[%v]"
)
