package main

import (
	"context"
	"go-framework/http/config"
	"go-framework/http/internal/server"
	"go-framework/http/pkg/logger"
	"go-framework/http/pkg/mysql"
	_ "go.uber.org/automaxprocs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var err error
	err = config.Init()
	if err != nil {
		panic("config.Init error:" + err.Error())
	}

	err = logger.InitBusinessLogger(&config.Config.Log)
	if err != nil {
		panic("logger.Init error:" + err.Error())
	}

	logger.Business.Logger.Info("哈喽～")

	err = mysql.NewMySqlPool(&mysql.Config{
		Master: config.Config.MysqlMaster,
		Slave:  config.Config.MysqlSlave,
	})
	if err != nil {
		panic("mysql.Init error:" + err.Error())
	}

	graceShutDown(server.NewHttpServer())
}

func graceShutDown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	<-quit

	log.Println("Shutdown Server ...")

	defer mysql.CloseMysqlPool() //close Mysql Pool

	//创建一个上下文, 10s后超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
