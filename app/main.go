package main

import (
	"flag"
	"fmt"

	"github.com/yrjkqq/tiny-website/app/build"
)

// Version  go build -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/yrjkqq/tiny-website/app/build.Time=$(date)' -X 'github.com/yrjkqq/tiny-website/app/build.User=$(id -u -n)'"
var Version = "dev"

// go run -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/yrjkqq/tiny-website/app/build.Time=$(date)' -X 'github.com/yrjkqq/tiny-website/app/build.User=$(id -u -n)'" main.go -version
// ./app -version
var showVersion = flag.Bool("version", false, "Print version of the binary")

func main() {
	fmt.Println("starting...")

	if !flag.Parsed() {
		flag.Parse()
	}
	fmt.Println("Version:\t", Version)
	if *showVersion {
		fmt.Println("build.Time:\t", build.Time)
		fmt.Println("build.User:\t", build.User)
		return
	}

	// demo here
	// demo.BeegoValidationDemo()
	// demo.GoUUIDDemo()
	// demo.GoRedisExampleClient()

	// gredis.Setup()

	// r := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        r,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// go func() {
	// 	if err := s.ListenAndServe(); err != nil {
	// 		log.Printf("Listen: %s\n", err)
	// 	}
	// }()

	// quit := make(chan os.Signal)
	// signal.Notify(quit, os.Interrupt)
	// <-quit

	// log.Println("Shutdown Server ...")

	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// if err := s.Shutdown(ctx); err != nil {
	// 	log.Fatal("Server Shutdown: ", err)
	// }

	// log.Println("Server exiting")

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
