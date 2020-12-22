package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-framework/http/config"
	"go-framework/http/internal/router"
	_ "go.uber.org/automaxprocs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	
	err := config.Init()
	if err != nil {
		panic("config.Init error:" + err.Error())
	}

	gin.SetMode(config.GetMode())
	g := gin.New()

	err = router.Init(g)
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

	graceShutDown(srv)
}

func graceShutDown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit

	log.Println("Shutdown Server ...")

	/*defer mysql.CloseMysqlPool()    //close Mysql Pool
	defer redis.CloseRedisPool()    //close Redis Pool*/

	//创建一个上下文, 10s后超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
