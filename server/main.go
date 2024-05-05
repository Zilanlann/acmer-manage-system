package main

import (
	"context"
	"fmt"
	"log"
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
	"github.com/zilanlann/acmer-manage-system/server/setting"
	_ "go.uber.org/automaxprocs"
)

func init() {
	setting.Setup()
	global.LOG = zap.Zap() // 初始化zap日志库
	model.Setup()
	redis.Setup()
}

func test() {
	contest := model.Contest{Name: "test", Desc: "test contest", StartTime: time.Now(), EndTime: time.Now(), Teams: []model.Team{{ZhName: "test team", CoachID: 73, Contestants: []model.Contestant{{UserID: 71}, {UserID: 14}, {UserID: 72}}}}}
	if err := model.CreateContest(&contest); err != nil {
		global.LOG.Error(err.Error())
	}
}

func main() {
	// test()
	// contest, _ := model.GetContestInfo(1)
	// json, _ := json.Marshal(contest)
	// fmt.Printf("json: %s\n", json)
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
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
