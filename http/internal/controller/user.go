package controller

import (
	"github.com/gin-gonic/gin"
	"go-framework/http/internal/dao"
	"go-framework/http/internal/protocol"
	"go-framework/http/internal/service"
	"go-framework/http/pkg/http"
)

func GetUser(c *gin.Context) {
	var req protocol.ReqGetUsers
	if err := http.GetBodyParam(c, &req); err != nil {
		http.ResponseError(c, err)
		return
	}

	//todo
	obj := service.NewUserService(dao.NewUserDao())
	resp := obj.GetUser()

	http.ResponseSuccess(c, resp)
	return
}
