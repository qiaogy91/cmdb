package main

import (
	_ "github.com/qiaogy91/cmdb/apps"
	_ "github.com/qiaogy91/ioc/apps/swagger/restful"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/ioc/server"
	_ "github.com/qiaogy91/maudit/provider/middleware"      // 审计
	_ "github.com/qiaogy91/mcenter/provider/authentication" // 认证
	_ "github.com/qiaogy91/mcenter/provider/endpoint"       // 注册api端点
	_ "github.com/qiaogy91/mcenter/provider/permission"     // 权限
)

func main() {
	server.Execute()
}
