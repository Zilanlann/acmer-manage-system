package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zilanlann/acmer-manage-system/server/global"
	"github.com/zilanlann/acmer-manage-system/server/model"
	"github.com/zilanlann/acmer-manage-system/server/pkg/redis"
	"github.com/zilanlann/acmer-manage-system/server/pkg/zap"
	"github.com/zilanlann/acmer-manage-system/server/routers"
	"github.com/zilanlann/acmer-manage-system/server/schedule"
	"github.com/zilanlann/acmer-manage-system/server/setting"
	_ "go.uber.org/automaxprocs"
)

func init() {
	setting.Setup()        // 初始化配置
	global.LOG = zap.Zap() // 初始化zap日志库
	model.Setup()          // 初始化数据库
	redis.Setup()          // 初始化redis
}

func main() {
	// 启动定时任务调度器
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go schedule.StartScheduler(ctx)

	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)

	s := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			global.LOG.Warn(err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	global.LOG.Info("Shutdown Server ...")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()
	if err := s.Shutdown(ctx2); err != nil {
		global.LOG.Warn(err.Error())
	}

	global.LOG.Info("Server exiting")
}
