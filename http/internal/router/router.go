package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-framework/http/internal/controller"
)

func Init(g *gin.Engine) (err error) {
	if g == nil {
		err = fmt.Errorf("nil gin engine")
		return err
	}

	//g.Use(middleware.Logger(), middleware.Recover())

	//微信业务类接口
	defaultGroup := g.Group("/")
	{
		defaultGroup.POST("/", controller.GetUser)
	}

	return nil
}
