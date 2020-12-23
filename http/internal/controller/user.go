package controller

import (
	"github.com/gin-gonic/gin"
	"go-framework/http/api"
	"go-framework/http/internal/di"
	"go-framework/http/pkg/http"
)

func GetUser(c *gin.Context) {
	var req api.ReqGetUsers
	if err := http.GetBodyParam(c, &req); err != nil {
		http.ResponseError(c, err)
		return
	}

	resp := di.NewUserService().GetUser(req.Uid)
	http.ResponseSuccess(c, resp)
	return
}
