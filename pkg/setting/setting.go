package setting

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Setting ...
var (
	Cfg *viper.Viper

	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	AccessJwtSecret     string
	RefreshJwtSecret    string
	AccessTokenExpires  time.Duration
	RefreshTokenExpires time.Duration

	PageSize int

	DBType        string
	DBUser        string
	DBPassword    string
	DBHost        string
	DBName        string
	DBTablePrefix string

	RedisCfg map[string]interface{}
)

func init() {
	viper.SetConfigName("app")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	Cfg = viper.GetViper()
	loadBase()
	loadServer()
	loadApp()
	loadDB()
	loadRedisCfg()
}

func loadBase() {
	Cfg.SetDefault("RunMode", "debug")
	RunMode = Cfg.GetString("RunMode")
}

func loadServer() {
	HTTPPort = Cfg.GetInt("HTTPPort")
	ReadTimeout = time.Duration(Cfg.GetInt("ReadTimeout")) * time.Second
	WriteTimeout = time.Duration(Cfg.GetInt("WriteTimeout")) * time.Second
}

func loadApp() {
	AccessJwtSecret = Cfg.GetString("AccessJwtSecret")
	RefreshJwtSecret = Cfg.GetString("RefreshJwtSecret")
	AccessTokenExpires = time.Duration(Cfg.GetInt("AccessTokenExpires"))
	RefreshTokenExpires = time.Duration(Cfg.GetInt("AccessTokenExpires"))
	PageSize = Cfg.GetInt("PageSize")
}

func loadDB() {
	DBType = Cfg.GetString("DBType")
	DBUser = Cfg.GetString("DBUser")
	DBPassword = Cfg.GetString("DBPassword")
	DBHost = Cfg.GetString("DBHost")
	DBName = Cfg.GetString("DBName")
	DBTablePrefix = Cfg.GetString("DBTablePrefix")
}

func loadRedisCfg() {
	// Important: Viper configuration keys are case insensitive. There are ongoing discussions about making that optional.
	RedisCfg = Cfg.GetStringMap("RedisCfg")
}
