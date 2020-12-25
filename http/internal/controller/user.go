package controller

import (
	"github.com/gin-gonic/gin"
	"go-framework/http/api"
	"go-framework/http/internal/di"
	"go-framework/http/pkg/http"
	"go-framework/http/pkg/logger"
)

func GetUser(c *gin.Context) {
	var req api.ReqGetUsers
	if err := http.GetBodyParam(c, &req); err != nil {
		http.ResponseError(c, err)
		return
	}

	resp, err := di.NewUserService().GetUser(req.Uid)
	if err != nil {
		logger.Warnf("数据问题 %+v", err)
		http.ResponseError(c, err)
		return
	}
	http.ResponseSuccess(c, resp)
	return
}
