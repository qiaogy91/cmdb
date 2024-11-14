package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrSecretParams(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10000, "参数不正确", e)
}
func ErrSecretUpgrade(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "协议升级失败", e)
}

func ErrSecretSync(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusInternalServerError, 10002, "同步过程出错", e)
}
