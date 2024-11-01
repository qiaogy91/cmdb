package main

import (
	_ "github.com/qiaogy91/cmdb/apps"
	_ "github.com/qiaogy91/ioc/config/otlp"
	"github.com/qiaogy91/ioc/server"
	_ "github.com/qiaogy91/maudit/provider/producer"
	_ "github.com/qiaogy91/mcenter/provider/middleware"
)

func main() {
	server.Execute()
}
