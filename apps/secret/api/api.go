package api

import (
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/gorilla/websocket"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"log/slog"
	"net/http"
)

type Handler struct {
	ioc.ObjectImpl
	log     *slog.Logger
	upgrade *websocket.Upgrader
	svc     secret.Service
}

func (h *Handler) Name() string  { return secret.AppName }
func (h *Handler) Priority() int { return 401 }

func (h *Handler) Init() {
	h.log = log.Sub(secret.AppName)
	h.svc = secret.GetSvc()
	h.upgrade = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
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
		Metadata("perm", true).
		Metadata("resource", secret.AppName).
		Metadata("action", "get"),
	)

	ws.Route(ws.GET("/SyncResource").To(h.SyncResource).
		Param(ws.QueryParameter("vendor", "厂商类型").DataType("string")).
		Param(ws.QueryParameter("secretId", "秘钥ID")).
		Param(ws.QueryParameter("accessEndpoint", "服务地址")).
		Param(ws.QueryParameter("regionId", "区域ID")).
		Doc("数据同步").
		Metadata(restfulspec.KeyOpenAPITags, tag).
		Metadata("auth", false).
		Metadata("audit", true).
		Metadata("perm", false).
		Metadata("resource", secret.AppName).
		Metadata("action", "sync"),
	)
}

func init() {
	ioc.Api().Registry(&Handler{})
}
