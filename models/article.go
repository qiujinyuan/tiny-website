package models

import (
	"log"

	"github.com/yrjkqq/tiny-website/pkg/logging"
)

// Article 文章
type Article struct {
	Base

	TagID string `json:"tagID" gorm:"index"`
	Tag   Tag    `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
	State      int    `json:"state"`
}

// ExistArticleByID 通过 id 判断 article 是否存在
func ExistArticleByID(id string) (exist bool, article Article) {
	if db.Where("id = ?", id).Find(&article).RecordNotFound() {
		exist = false
		return
	}
	exist = true
	return
}

// GetArticleTotal article 总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

// GetArticles article 列表
func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//db.Model(&user).Related(&profile)
//// SELECT * FROM profiles WHERE user_id = 111; // 111 is user's ID

// GetArticle article 详情
func GetArticle(id string) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

// EditArticle edit an article
func EditArticle(id string, data interface{}) bool {
	err := db.Model(&Article{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		log.Println("Edit article failed: ", id, err)
		return false
	}
	return true
}

// AddArticle add an article
func AddArticle(data map[string]interface{}) bool {
	tagID, ok := data["tagID"].(string)
	if !ok {
		tagID = ""
	}
	err := db.Create(&Article{
		TagID:     tagID,
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["createdBy"].(string),
		State:     data["state"].(int),
	}).Error

	if err != nil {
		log.Println("Unable to create article: ", data, err)
		return false
	}
	return true
}

// DeleteArticle delete the article
func DeleteArticle(id string) (err error) {
	err = db.Where("id = ?", id).Delete(&Article{}).Error
	if err != nil {
		log.Println("DeleteArticle failed: ", id, err)
		return
	}
	return
}

// ClearAllSoftDeletedArticle 删除所有已经软删除的 article
func ClearAllSoftDeletedArticle() bool {
	err := db.Unscoped().Where("deleted_at IS NOT NULL").Delete(&Article{}).Error
	if err != nil {
		logging.Error("Clear all soft deleted article failed: ", err)
		return false
	}
	return true
}
