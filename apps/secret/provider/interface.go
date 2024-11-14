package provider

import (
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc/config/log"
	"log/slog"
)

const AppName = "secretProvider"

type Provider interface {
	Name() string
	Type() secret.Vendor
	Init(Conf)

	SyncResource(request *secret.SyncResourceRequest, server secret.Rpc_SyncResourceServer) error
}

var (
	//containerLogger =
	container   = map[secret.Vendor]Provider{}
	AddProvider = func(o Provider) { container[o.Type()] = o }
	GetProvider = func(t secret.Vendor) Provider { return container[t] }

	InitProvider = func(conf Conf) {
		l := log.Sub(AppName)
		for _, o := range container {
			o.Init(conf)
			l.Info("add provider", slog.String("name", o.Name()), slog.Any("type", o.Type()))
		}
	}
)
