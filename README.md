# Tiny website

## 网站配色方案

[Ant Design 基础色板-明青](https://ant.design/docs/spec/colors-cn#%E5%9F%BA%E7%A1%80%E8%89%B2%E6%9D%BF)

- #e6fffb
- #b5f5ec
- #87e8de
- #5cdbd3
- #36cfc9 font-color: rgba(0, 0, 0, 0.85);
- #13c2c2 font-color: rgb(255, 255, 255);
- #08979c
- #006d75
- #00474f
- #002329

## markdown 编辑器

使用 marked.js 实现，左侧输入右侧预览，参考 [Marked Demo](https://marked.js.org/demo)

增加上传图片按钮，上传至七牛云图床，返回图片图片地址，以 ```![图片名称](图片url)``` 的格式 append 到输入框内

## 样式

尽量不把样式抽离到 css 中，使用 ```{{ define "css-common" }}``` 定义通用样式并在其他模版中使用 ```{{template "css-common" .}}``` 的方式引入，可以将样式按照 style 的方式加载，那么下载下来的 html 文件中已经有样式

## 学习资料

- [跟煎鱼学 Go](https://eddycjy.com/go-categories/)
- [Build Web Application with Golang](https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/preface.md)
- [JWT introduction](https://jwt.io/introduction/)
- [Using JWT for Authentication in a Golang Application](https://www.nexmo.com/blog/2020/03/13/using-jwt-for-authentication-in-a-golang-application-dr)
- [Debugging a Go Web App with VSCode and Delve](https://www.thegreatcodeadventure.com/debugging-a-go-web-app-with-vscode-and-delve/)
- [Debugging vscode-go](https://github.com/golang/vscode-go/blob/master/docs/debugging.md)
- [How to Use UUID Key Type with Gorm](https://medium.com/@the.hasham.ali/how-to-use-uuid-key-type-with-gorm-cc00d4ec7100)
- [Store UUID v4 in MySQL](https://stackoverflow.com/questions/43056220/store-uuid-v4-in-mysql)
- [Storing UUID Values in MySQL Tables](https://mysqlserverteam.com/storing-uuid-values-in-mysql-tables/)
- [如何重构“箭头型”代码](https://coolshell.cn/articles/17757.html)
- [Go Concurrency Patterns: Context](https://blog.golang.org/context)
- [The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets (No Excuses!)](https://www.joelonsoftware.com/2003/10/08/the-absolute-minimum-every-software-developer-absolutely-positively-must-know-about-unicode-and-character-sets-no-excuses/)

## 库

- [jwt-go](https://github.com/dgrijalva/jwt-go)
- [beego-validation](https://beego.me/docs/mvc/controller/validation.md)
- [gorm](https://gorm.io/docs/query.html#Where)
- [gin](https://gin-gonic.com/docs/examples/controlling-log-output-coloring/)
- [UUID package for Go language](https://github.com/satori/go.uuid)
- [go-redis](https://github.com/go-redis/redis)

## swagger

- [swagger](http://127.0.0.1:10086/swagger/index.html)