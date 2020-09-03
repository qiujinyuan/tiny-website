package main

import (
	"flag"

	"github.com/yrjkqq/tiny-website/demo"
)

// Version  go build -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/yrjkqq/tiny-website/app/build.Time=$(date)' -X 'github.com/yrjkqq/tiny-website/app/build.User=$(id -u -n)'"
var Version = "dev"

// go run -ldflags="-X 'main.Version=v1.0.0' -X 'github.com/yrjkqq/tiny-website/app/build.Time=$(date)' -X 'github.com/yrjkqq/tiny-website/app/build.User=$(id -u -n)'" main.go -version
// ./app -version
var showVersion = flag.Bool("version", false, "Print version of the binary")

var execDemo = flag.Bool("demo", false, "Exec demo")

func main() {
	if !flag.Parsed() {
		flag.Parse()
	}

	if *execDemo {
		// demo here
		// demo.BeegoValidationDemo()
		// demo.GoUUIDDemo()
		// demo.GoRedisExampleClient()
		// demo.ContextPkgUseDemo()
		// demo.CronUseDemo()
		demo.GoqueryExample()
		return
	}

	// fmt.Println("Version:\t", Version)
	// if *showVersion {
	// 	fmt.Println("build.Time:\t", build.Time)
	// 	fmt.Println("build.User:\t", build.User)
	// 	return
	// }

	// go app.ServeBackground()

	// app.Start()
}
