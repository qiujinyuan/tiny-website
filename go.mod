module github.com/yrjkqq/tiny-website

go 1.14

require (
	github.com/360EntSecGroup-Skylar/excelize v1.3.1-0.20180527032555-9e463b461434
	github.com/EDDYCJY/go-gin-example v0.0.0-20200505102242-63963976dee0
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/astaxie/beego v1.12.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-redis/redis/v8 v8.0.0-beta.5
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/microcosm-cc/bluemonday v1.0.3
	github.com/robfig/cron/v3 v3.0.1
	github.com/satori/go.uuid v1.2.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
	github.com/swaggo/files v0.0.0-20190704085106-630677cd5c14
	github.com/swaggo/gin-swagger v1.2.0
	github.com/swaggo/swag v1.5.1
	github.com/tealeg/xlsx v1.0.4-0.20180419195153-f36fa3be8893
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20200707034311-ab3426394381 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
