package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/gorilla/websocket"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc"
	"github.com/qiaogy91/ioc/config/gorestful"
	"github.com/qiaogy91/ioc/config/log"
	"github.com/qiaogy91/ioc/labels"
	"github.com/qiaogy91/ioc/utils"
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

type User struct {
	Name string
	Age  int
}

type UserSet struct {
	total int
	items []*User
}

func (h *Handler) Registry() {
	tag := []string{"凭证管理"}

	ws := gorestful.ModuleWebservice(h)
	ws.Route(ws.GET("/SyncResource").To(h.SyncResource).
		Param(ws.QueryParameter("vendor", "厂商类型").DataType("string")).
		Param(ws.QueryParameter("secretId", "秘钥ID")).
		Param(ws.QueryParameter("accessEndpoint", "服务地址")).
		Param(ws.QueryParameter("regionId", "区域ID")).
		Doc("数据同步").
		Metadata(labels.ApiTags, tag).
		Metadata(labels.AuthEnabled, false).
		Metadata(labels.AuditEnabled, true).
		Metadata(labels.PermissionEnabled, false).
		Metadata(labels.Resource, secret.AppName).
		Metadata(labels.Action, "sync").
		Metadata(labels.Desc, "资源同步"),
	)

	ws.Route(ws.GET("test").To(func(r *restful.Request, w *restful.Response) {
		utils.SendSuccess(w, "ok")
	}).
		Doc("测试一下子").
		Metadata(labels.ApiTags, tag).
		Metadata(labels.AuthEnabled, true).
		Metadata(labels.AuditEnabled, true).
		Metadata(labels.PermissionEnabled, true).
		Metadata(labels.Resource, secret.AppName).
		Metadata(labels.Action, "get").
		Metadata(labels.Desc, "测试资源"),
	)
}

func init() {
	ioc.Api().Registry(&Handler{})
}
