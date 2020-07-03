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

	JwtSecret string

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
	JwtSecret = Cfg.GetString("JwtSecret")
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
	RedisCfg = Cfg.GetStringMap("RedisCfg")
}
