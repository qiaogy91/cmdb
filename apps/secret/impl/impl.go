package impl

import (
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/cmdb/apps/secret/provider"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/datasource"
	"github.com/qiaogy91/ioc/config/grpc"
	"github.com/qiaogy91/ioc/config/log"
	"gorm.io/gorm"
	"log/slog"
)

var _ secret.RpcServer = &Impl{}

type Impl struct {
	ioc.ObjectImpl
	secret.UnimplementedRpcServer
	log      *slog.Logger
	db       *gorm.DB
	Provider provider.Conf `json:"provider" yaml:"provider"`
}

func (i *Impl) Name() string  { return secret.AppName }
func (i *Impl) Priority() int { return 301 }
func (i *Impl) Init() {
	i.log = log.Sub(secret.AppName)
	i.db = datasource.DB()

	provider.InitProvider(i.Provider) // 初始化provider
	i.gRpcRegistry()                  // 注册grpc 服务
}

func (i *Impl) gRpcRegistry() {
	s := grpc.Get().Server()
	secret.RegisterRpcServer(s, i)
}

func init() {
	ioc.Controller().Registry(&Impl{})
}
