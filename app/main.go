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
	fmt.Println("start...")
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

	// fmt.Printf("%q\n", time.Now().Format("2006/01/02 - 15:04:05.00000"))
	// fmt.Printf("%q\n", time.Now().Format("2006/01/02 - 15:04:05.99999999"))

	// gredis.Setup()

	// r := routers.InitRouter()

	// s := &http.Server{
	// 	Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
	// 	Handler:        r,
	// 	ReadTimeout:    setting.ReadTimeout,
	// 	WriteTimeout:   setting.WriteTimeout,
	// 	MaxHeaderBytes: 1 << 20,
	// }

	// s.ListenAndServe()
}
