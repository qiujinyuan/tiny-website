package routers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/yrjkqq/tiny-website/middleware/jwt"
	"github.com/yrjkqq/tiny-website/pkg/setting"
	"github.com/yrjkqq/tiny-website/routers/api"
	v1 "github.com/yrjkqq/tiny-website/routers/api/v1"
	"github.com/yrjkqq/tiny-website/testdata/protoexample"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Post blog object
type Post struct {
	ID      string `uri:"id" binding:"required"`
	Content string
}

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	gin.SetMode(setting.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", api.GetAuth)

	r.GET("/logout", jwt.JWT(), api.Logout)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)

		//获取文章列表
		apiV1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiV1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiV1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiV1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	// 加载静态资源
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")
	r.Static("/static", "./static")

	// 加载 html 模板
	r.LoadHTMLGlob("templates/**/*")

	r.GET("/ping", func(c *gin.Context) {
		// c.JSON(200, gin.H{
		// 	"message": "pong",
		// })
		// c.HTML(http.StatusUnauthorized, name string, obj interface{})
		fmt.Println(c.Request.Header.Get("Authorization"))
		c.Header("WWW-Authenticate", "fuck off.")
		// c.JSON(http.StatusUnauthorized, gin.H{
		// 	"message": "401",
		// })
		c.JSON(http.StatusUnauthorized, nil)
	})

	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	r.GET("/posts/:id", func(c *gin.Context) {
		var post Post
		uriError := c.ShouldBindUri(&post)
		if uriError != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"title": "not found",
			})
			return
		}
		mdContent, readErr := ioutil.ReadFile("./resources/posts/" + post.ID + ".md")
		if readErr != nil {
			c.HTML(http.StatusNotFound, "404.html", gin.H{
				"title": post.ID + ".md not found",
			})
			return
		}
		unsafe := blackfriday.Run(mdContent)
		// 从客户端获取到 markdown 内容后需要向使用以下代码进行数据清洗
		safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": post.ID,
			// 告诉模版 content 是 html 片段，不需要转义
			"content": template.HTML(string(safe)),
		})
	})

	// 新建文章
	r.GET("/post/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/new.html", gin.H{})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
			"title": "<Users />",
		})
	})

	return r
}
