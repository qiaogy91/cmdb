package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"log/slog"
)

type Handler struct {
	ioc.ObjectImpl
	log *slog.Logger
}

func (h *Handler) Name() string  { return secret.AppName }
func (h *Handler) Priority() int { return 401 }

func (h *Handler) Init() {
	h.log = log.Sub(secret.AppName)
	h.Registry()
}

func (h *Handler) Registry() {
	tag := []string{"凭证管理"}

	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.GET("/").To(h.Test).
		Doc("测试接口").
		Metadata(restfulspec.KeyOpenAPITags, tag).
		Metadata("auth", true).
		Metadata("audit", true).
		Metadata("resource", secret.AppName).
		Metadata("action", "get"),
	)
}

func init() {
	ioc.Api().Registry(&Handler{})
}
