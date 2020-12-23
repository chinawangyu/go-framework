package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-framework/http/config"
	"go-framework/http/internal/controller"
	"log"
	"net/http"
	"time"
)

//启动http服务
func NewHttpServer() *http.Server {
	gin.SetMode(config.GetMode())
	g := gin.New()

	err := routerInit(g)
	if err != nil {
		panic("router.Init error:" + err.Error())
	}

	port := config.GetPort()
	srv := &http.Server{
		Addr:           port,
		Handler:        g,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen port %s error", port)
		}
	}()

	return srv
}

//路由
func routerInit(g *gin.Engine) (err error) {
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
