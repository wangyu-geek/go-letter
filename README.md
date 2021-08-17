
## go-leeter
基于beego的go服务端脚手架，快速构建go项目

## 目录结构

```
├── controllers 控制器
├──conf        配置文件
|  └── app.conf
├── docs       文档
├── logs       业务日志
├── models     模型
├── main.go    入口文件
├── pkg        封装一些常用类库
|  ├── logging 日志工具类
|  ├── setting 系统配置
├── routers    定义路由
├── serializer 统一处理接口的返回值
├── service    处理复杂业务
├── middlewares 中间件
```

## 构建脚本
`./build-linux.sh`


## 项目运行
windiws：`go run main.go`
linux：`./leeter.linux`

## 配置文件
[/conf/app.conf](./conf/app.conf)


## 接口
1. ping

url：http://127.0.0.1:端口号/ping

