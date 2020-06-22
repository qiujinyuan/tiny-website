package setting

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	// HTTPPort ...
	HTTPPort int
	// ReadTimeout ...
	ReadTimeout time.Duration
	// WriteTimeout ...
	WriteTimeout time.Duration
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	loadServer()
}

func loadServer() {
	HTTPPort = viper.GetInt("http_port")
	ReadTimeout = time.Duration(viper.GetInt("read_timeout")) * time.Second
	WriteTimeout = time.Duration(viper.GetInt("write_timeout")) * time.Second
}
