package v1

import (
	"fmt"
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
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
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("createdBy")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 100, "name").Message("名称最长为 100 字符")
	valid.MinSize(name, 3, "name").Message("名称最短为 3 字符")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "createdBy").Message("创建人最长为 100 字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许 0 或 1")

	code := e.InvalidParams
	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		if models.ExistTagByName(name) {
			code = e.ErrorExistTag
		} else {
			models.AddTag(name, state, createdBy)
			code = e.Success
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
		"data": make(map[string]interface{}),
	})

}

// EditTag 修改文章标签
func EditTag(c *gin.Context) {}

// DeleteTag 删除文章标签
func DeleteTag(c *gin.Context) {}
