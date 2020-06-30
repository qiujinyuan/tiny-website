package models

import (
	"log"
)

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
func ExistTagByName(name string) (exist bool, tag Tag) {
	if db.Where("name = ?", name).Find(&tag).RecordNotFound() {
		exist = false
		return
	}
	exist = true
	return
}

// ExistTagById 通过 id 判断 tag 是否存在
func ExistTagById(id string) (exist bool, tag Tag) {
	if db.Where("id = ?", id).Find(&tag).RecordNotFound() {
		exist = false
		return
	}
	exist = true
	return
}

// AddTag 增加 tag
func AddTag(name string, state int, createdBy string) bool {
	err := db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
	if err != nil {
		log.Println("Unable to create user: ", err)
		return false
	}
	return true
}

// DeleteTag 删除标签
func DeleteTag(id string) bool {
	err := db.Where("id = ?", id).Delete(&Tag{}).Error
	if err != nil {
		log.Println("Delete tag by id failed: ", id, err)
		return false
	}
	return true
}

// 编辑标签
func EditTag(id string, data interface{}) bool {
	err := db.Model(&Tag{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		log.Println("Edit tag failed: ", id, err)
		return false
	}
	return true
}
