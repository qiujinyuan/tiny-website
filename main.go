package main

import (
	"fmt"
	"net/http"

	"github.com/yrjkqq/tiny-website/pkg/gredis"
	"github.com/yrjkqq/tiny-website/pkg/setting"
	"github.com/yrjkqq/tiny-website/routers"
)

func main() {
	fmt.Println("starting...")

	// demo here
	// demo.BeegoValidationDemo()
	// demo.GoUUIDDemo()
	// demo.GoRedisExampleClient()

	gredis.Setup()

	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
