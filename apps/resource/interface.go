package resource

import (
	"context"
	"github.com/qiaogy91/ioc"
)

const AppName = "resource"

func GetSvc() Service { return ioc.Controller().Get(AppName).(Service) }

type Service interface {
	RpcServer
	CreateTable(ctx context.Context) error
}
