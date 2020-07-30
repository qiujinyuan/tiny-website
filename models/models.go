package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"github.com/yrjkqq/tiny-website/pkg/setting"

	// mysql
	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

// Model ...
type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"createdOn"`
	ModifiedOn int `json:"modifiedOn"`
}

// Base contains common columns for all tables.
type Base struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CreatedAt  int       `json:"createdAt"`
	ModifiedAt int       `json:"modifiedAt"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	u1 := uuid.NewV4()
	scope.SetColumn("ID", u1)
	scope.SetColumn("CreatedAt", time.Now().Unix())

	return nil
}

// BeforeUpdate gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。
func (base *Base) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedAt", time.Now().Unix())

	return nil
}

func init() {
	var err error
	db, err = gorm.Open(setting.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", setting.DBUser, setting.DBPassword, setting.DBHost, setting.DBName))
	if err != nil {
		log.Fatalln(err)
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
