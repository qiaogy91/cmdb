package ali

import (
	"github.com/qiaogy91/cmdb/apps/resource"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/cmdb/apps/secret/provider"
	"github.com/qiaogy91/ioc/config/datasource"
	"gorm.io/gorm"
	"log/slog"
)

const AppName = "ali"

var _ provider.Provider = &Impl{}

type Impl struct {
	log      *slog.Logger
	db       *gorm.DB
	resource resource.Service
}

func (i *Impl) Name() string        { return AppName }
func (i *Impl) Type() secret.Vendor { return secret.Vendor_VENDOR_ALI }
func (i *Impl) Init(conf provider.Conf) {
	i.db = datasource.DB()
	i.resource = resource.GetSvc()
}

func init() {
	provider.AddProvider(&Impl{})
}
