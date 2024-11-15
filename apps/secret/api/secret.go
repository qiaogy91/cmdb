package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc/utils"
)

func (h *Handler) SyncResource(r *restful.Request, w *restful.Response) {
	// 构建请求
	req := &secret.SyncResourceRequest{}
	if err := utils.Decoder.Decode(req, r.Request.URL.Query()); err != nil {
		utils.SendFailed(w, ErrSecretParams(err))
		return
	}

	// 升级协议
	conn, err := h.upgrade.Upgrade(w, r.Request, nil)
	if err != nil {
		utils.SendFailed(w, ErrSecretUpgrade(err))
		return
	}

	// 执行同步
	defer func() { _ = conn.Close() }()
	if err := h.svc.SyncResource(req, NewStream(conn)); err != nil {
		utils.WebsocketSendFailed(conn, ErrSecretSync(err))
		return
	}

	// 结束标识
	utils.WebsocketSendSuccess(conn, "资源同步结束")
}
