module github.com/yrjkqq/tiny-website

go 1.14

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/go-playground/validator/v10 v10.3.0 // indirect
	github.com/golang/protobuf v1.4.2
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/microcosm-cc/bluemonday v1.0.3
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/spf13/viper v1.7.0
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/sys v0.0.0-20200615200032-f1bc736245b1 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/russross/blackfriday.v2 v2.0.1
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace (
	github.com/yrjkqq/tiny-website/pkg/setting => ./pkg/setting
	github.com/yrjkqq/tiny-website/testdata/protoexample => ./testdata/protoexample
	gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
)
