package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yrjkqq/tiny-website/pkg/setting"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// Model ...
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"createdOn"`
	ModifiedOn int `json:"modifiedOn"`
}

func init() {
	var err error
	db, err = gorm.Open(setting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DBUser, setting.DBPassword, setting.DBHost, setting.DBName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DBTablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	defer db.Close()
}
