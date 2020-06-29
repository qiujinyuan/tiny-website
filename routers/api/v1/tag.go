package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/yrjkqq/tiny-website/models"
	"github.com/yrjkqq/tiny-website/pkg/e"
	"github.com/yrjkqq/tiny-website/pkg/setting"
	"github.com/yrjkqq/tiny-website/pkg/util"
)

// GetTags 获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int = -1

	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	code := e.Success
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})

}

// AddTag 新增文章标签
func AddTag(c *gin.Context) {}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {}
