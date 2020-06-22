package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/qiujinyuan/tiny-website/testdata/protoexample"

	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Post blog object
type Post struct {
	ID      string `uri:"id" binding:"required"`
	Content string
}

func main() {
	fmt.Println("starting...")
	r := gin.Default()

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
		c.Header("WWW-Authenticate", "fuck off")
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

	r.Run()
}
