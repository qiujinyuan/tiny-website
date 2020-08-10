package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yrjkqq/tiny-website/models"
	"github.com/yrjkqq/tiny-website/pkg/gredis"
	"github.com/yrjkqq/tiny-website/pkg/logging"
	"github.com/yrjkqq/tiny-website/pkg/setting"
	"github.com/yrjkqq/tiny-website/routers"
)

// Start start the tiny-website
func Start() {
	fmt.Println("starting...")
	setting.Setup()
	models.Setup()
	logging.Setup()

	gredis.Setup()

	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout * time.Second,
		WriteTimeout:   setting.ServerSetting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown: ", err)
	}

	log.Println("Server exiting")

	// endless 在 windows 下面无法使用, 有几个信号 windows 下未定义
	// endless.DefaultReadTimeOut = setting.ReadTimeout
	// endless.DefaultWriteTimeOut = setting.WriteTimeout
	// endless.DefaultMaxHeaderBytes = 1 << 20
	// endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Serve err: %v", err)
	// }
}
