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

## 构建

使用 ldflags 设置构建参数时, 除 main 包以外需要使用完整路径包名

- [Using ldflags to Set Version Information for Go Applications](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications)

## swagger

- [swagger](http://127.0.0.1:10086/swagger/index.html)

## docker 

docker 配置国内镜像源
https://www.jianshu.com/p/405fe33b9032
阿里云
镜像加速器地址: https://lz7zbl35.mirror.aliyuncs.com
操作文档: https://cr.console.aliyun.com/cn-hangzhou/instances/mirrors

构建本地镜像
```bash
docker build -t tiny-website-docker .
```

本地镜像构建完成后, 启动 docker 容器
```bash
docker run --link mysql:mysql --link redis:redis -p 8000:10086 tiny-website-docker
```

增加命令 --link mysql:mysql 让 Golang 容器与 Mysql 容器互联；通过 --link，可以在容器内直接使用其关联的容器别名进行访问，而不通过 IP，但是--link只能解决单机容器间的关联，在分布式多机的情况下，需要通过别的方式进行连接

而且, --link 的方式要求配置文件的 mysql 的 host 必须为 mysql:3306, 而这样在本地是无法进行开发的

应该是使用 docker-compose 的方式来进行连接

docker 启动 mysql 容器中的 mysql 命令行
```bash
docker exec -it mysql mysql -uroot -p
```
or
```bash
docker exec -it mysql bash
```

## mysql

docker 安装 mysql
```bash
docker pull mysql
```

启动 mysql container
```bash
docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=rootroot -d mysql
```

挂载数据卷的方式启动 mysql, 重新创建容器之后恢复数据
```bash
docker run --name mysql -p 3307:3306 -e MYSQL_ROOT_PASSWORD=rootroot -v C:\data\docker-mysql:/var/lib/mysql -d mysql
```

windows 访问 docker 内安装的 mysql
```bash
docker exec -it mysql mysql -uroot -p
```
无法直接使用 mysql 命令直接访问

程序中应该使用 3307 端口进行访问

## redis
 
docker 安装 redis, 会自动拉取 redis 镜像
```
docker run --name redis -p 6379:6379 -v C:\data\docker-redis:/data -d redis redis-server --appendonly yes
docker exec -it redis bash
redis-cli
```
