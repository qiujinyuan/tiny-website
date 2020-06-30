module github.com/yrjkqq/tiny-website

go 1.14

require (
	github.com/astaxie/beego v1.12.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0
	github.com/golang/protobuf v1.4.2
	github.com/jinzhu/gorm v1.9.14
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/microcosm-cc/bluemonday v1.0.3
	github.com/satori/go.uuid v1.2.0
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
	github.com/unknwon/com v1.0.1
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

// github.com/yrjkqq/tiny-website/routers => ./routers
// github.com/yrjkqq/tiny-website/models => ./models
// github.com/yrjkqq/tiny-website/pkg/util => ./pkg/util
// github.com/yrjkqq/tiny-website/pkg/e => ./pkg/e
// github.com/yrjkqq/tiny-website/pkg/setting => ./pkg/setting
// github.com/yrjkqq/tiny-website/testdata/protoexample => ./testdata/protoexample
replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
