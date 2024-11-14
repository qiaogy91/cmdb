package impl

import (
	"github.com/qiaogy91/cmdb/apps/resource"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"gorm.io/gorm"
	"log/slog"
)

var _ resource.Service = &Impl{}

type Impl struct {
	resource.UnimplementedRpcServer
	ioc.ObjectImpl
	log *slog.Logger
	db  *gorm.DB
}

func (i *Impl) Name() string { return resource.AppName }

func (i *Impl) Priority() int { return 302 }

func (i *Impl) Init() {
	i.log = log.Sub(resource.AppName)
	i.db = datasource.DB()
}

func (i *Impl) gRpcRegistry() {
	s := grpc.Get().Server()
	resource.RegisterRpcServer(s, i)
}
func init() {
	ioc.Controller().Registry(&Impl{})
}
