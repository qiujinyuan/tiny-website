package setting

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type App struct {
	PageSize int

	AccessJwtSecret     string
	RefreshJwtSecret    string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration
}

type RedisCfg struct {
	Addr        string
	Network     string
	DialTimeout time.Duration
	Password    string

	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

type MySQLCfg struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var ServerSetting = &Server{}

var AppSetting = &App{}

var RedisSetting = &RedisCfg{}

var MySQLSetting = &MySQLCfg{}

// Setting ...
var cfg *viper.Viper

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	cfg = viper.GetViper()

	load("Server", ServerSetting)
	load("App", AppSetting)
	load("MySQL", MySQLSetting)
	load("RedisCfg", RedisSetting)
}

func load(subName string, targetCfg interface{}) {
	// Important: Viper configuration keys are case insensitive. There are ongoing discussions about making that optional.
	err := cfg.Sub(subName).Unmarshal(targetCfg)
	if err != nil {
		log.Fatalf("unable to decode into %v struct, %v", subName, err)
	}
}
