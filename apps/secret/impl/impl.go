package impl

import (
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/log"
	"gorm.io/gorm"
	"log/slog"
)

var _ secret.RpcServer = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	secret.UnimplementedRpcServer
	log       *slog.Logger
	db        *gorm.DB
	container ioc.ProviderContainer
}

func (i *Impl) Name() string  { return secret.AppName }
func (i *Impl) Priority() int { return 301 }
func (i *Impl) Init() {
	i.log = log.Sub(secret.AppName)
	i.db = datasource.DB()
	i.container = ioc.NewProviderContainer()

	i.providerRegistry()
}

func (i *Impl) providerRegistry() {
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
