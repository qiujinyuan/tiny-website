package v1

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/yrjkqq/tiny-website/models"
	"github.com/yrjkqq/tiny-website/pkg/e"
	"github.com/yrjkqq/tiny-website/pkg/logging"
	"github.com/yrjkqq/tiny-website/pkg/setting"
	"github.com/yrjkqq/tiny-website/pkg/util"
)

// GetArticle 获取单个文章
func GetArticle(c *gin.Context) {
	id := c.Param("id")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID 不能为空")

	code := e.INVALID_PARAMS
	var msg string
	var data interface{}
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		if exist, _ := models.ExistArticleByID(id); exist {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
		"data": data,
	})
}

// GetArticles 获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}

	var state int
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许 0 或 1")
	}

	tagID := c.Query("tagID")
	if tagID != "" {
		maps["tagID"] = tagID
	}

	code := e.INVALID_PARAMS
	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		code = e.SUCCESS
		data["list"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
		"data": data,
	})
}

// AddArticle 新增文章
// tagID 可以为空
func AddArticle(c *gin.Context) {
	tagID := c.Query("tagID")
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("createdBy")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "createdBy").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许 0 或 1")

	code := e.INVALID_PARAMS
	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		if tagID != "" {
			if exist, _ := models.ExistTagByID(tagID); exist {
				success := models.AddArticle(map[string]interface{}{
					"tagID":     tagID,
					"title":     title,
					"desc":      desc,
					"content":   content,
					"createdBy": createdBy,
					"state":     state,
				})
				if success {
					code = e.SUCCESS
				} else {
					code = e.ERROR
				}
			} else {
				code = e.ERROR_NOT_EXIST_TAG
			}
		} else {
			success := models.AddArticle(map[string]interface{}{
				"title":     title,
				"desc":      desc,
				"content":   content,
				"createdBy": createdBy,
				"state":     state,
			})
			if success {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
	})
}

// EditArticle 修改文章
// 修改时不允许清空 tagID
func EditArticle(c *gin.Context) {
	valid := validation.Validation{}

	id := c.Param("id")
	tagID := c.Query("tagID")
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modifiedBy")

	var state int
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许 0 或 1")
	}

	valid.Required(id, "id").Message("ID 不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modifiedBy").Message("修改人最长为100字符")

	code := e.INVALID_PARAMS
	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		if exist, _ := models.ExistArticleByID(id); exist {
			data := make(map[string]interface{})
			if tagID != "" {
				existTag, _ := models.ExistTagByID(tagID)
				if !existTag {
					code = e.ERROR_NOT_EXIST_TAG
					c.JSON(http.StatusOK, gin.H{
						"code": code,
						"msg":  e.GetMsg(code),
					})
					return
				}
				data["tagID"] = tagID
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			data["modifiedBy"] = modifiedBy
			success := models.EditArticle(id, data)
			if success {
				code = e.SUCCESS
			} else {
				code = e.ERROR
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
	})

}

// DeleteArticle 删除文章
func DeleteArticle(c *gin.Context) {
	id := c.Param("id")

	valid := validation.Validation{}
	valid.Required(id, "id").Message("ID 不能为空")

	code := e.INVALID_PARAMS
	var msg string
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			logging.Info(fmt.Sprintf("err.key: %s, err.message: %s", err.Key, err.Message))
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	} else {
		if exist, _ := models.ExistArticleByID(id); exist {
			err := models.DeleteArticle(id)
			if err != nil {
				code = e.ERROR
				msg += err.Error()
			} else {
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
	})
}
