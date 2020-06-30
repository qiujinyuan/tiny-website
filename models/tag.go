package models

import "log"

// Tag 文章标签
type Tag struct {
	Base

	Name       string `json:"name"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
	State      int    `json:"state"`
}

// GetTags 获取标签列表
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

// GetTagTotal 获取标签总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

// ExistTagByName 通过名称判断 tag 是否存在
func ExistTagByName(name string) bool {
	var tag Tag
	if !db.Where("name = ?", name).Find(&tag).RecordNotFound() {
		return true
	}
	return false
}

// AddTag 增加 tag
func AddTag(name string, state int, createdBy string) bool {
	err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
	if err != nil {
		log.Panic("Unable to create user", err)
	}
	return true
}
