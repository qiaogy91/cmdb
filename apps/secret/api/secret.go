package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/cmdb/apps/secret"
	"github.com/qiaogy91/ioc/utils"
	"log/slog"
	"strconv"
)

func (h *Handler) Test(r *restful.Request, w *restful.Response) {
	utils.SendSuccess(w, "ok")
	h.log.Info("test")
}

func (h *Handler) SyncResource(r *restful.Request, w *restful.Response) {
	// 获取请求参数 sync request
	vid, err := strconv.ParseInt(r.QueryParameter("vendor"), 10, 32)
	if err != nil {
		utils.SendFailed(w, ErrSecretParams(err))
		return
	}
	sid, err := strconv.ParseInt(r.QueryParameter("secretId"), 10, 64)
	if err != nil {
		utils.SendFailed(w, ErrSecretParams(err))
		return
	}

	// 构建请求
	req := &secret.SyncResourceRequest{
		Vendor:         secret.Vendor(vid),
		SecretId:       sid,
		AccessEndpoint: r.QueryParameter("accessEndpoint"),
		RegionId:       r.QueryParameter("regionId"),
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
	if err := conn.WriteJSON("资源同步结束"); err != nil {
		h.log.Error("send ws response err", slog.Any("err", err))
	}
}
