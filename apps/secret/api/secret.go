package api

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/qiaogy91/ioc/utils"
)

func (h *Handler) Test(r *restful.Request, w *restful.Response) {
	utils.SendSuccess(w, "ok")
	h.log.Info("test")
}
