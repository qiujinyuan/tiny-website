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
	CreatedAt  *int      `json:"createdAt"`
	ModifiedAt *int      `json:"modifiedAt"`
	// 如果一个 model 有 DeletedAt 字段，他将自动获得软删除的功能！ 当调用 Delete 方法时， 记录不会真正的从数据库中被删除， 只会将 DeletedAt 字段的值会被设置为当前时间
	// 查询语句为 SELECT * FROM `blog_tag`  WHERE `blog_tag`.`deleted_at` IS NULL AND ((name = '333'))，会自动拼接上 deleted_at is null, 导致数据库中 deleted_at 为 0 的记录查询不上来
	// 表结构需要修改为 `deleted_at` int(10) DEFAULT NULL,
	// 但是调用 Delete 时更新调用的 sql 是： UPDATE `blog_tag` SET `deleted_at`='2020-08-02 23:05:12'  WHERE `blog_tag`.`deleted_at` IS NULL AND ((id = '1ae27cc2-5f1f-463f-adf1-915274a7ad6d'));
	// 字符串的格式与数据库的格式不匹配，无法存储，报错：ERROR 1265 (01000): Data truncated for column 'deleted_at' at row 1
	// 如何解决？无法使用 gorm 提供的 BeforeDelete hooks 来指定时间。
	// 参考源码，使用自定义的 callback 替换 gorm 的 deleteCallback，-https://github.com/jinzhu/gorm/blob/master/callback_delete.go, -https://gorm.io/zh_CN/docs/write_plugins.html#%E9%A2%84%E5%AE%9A%E4%B9%89%E5%9B%9E%E8%B0%83
	// 成功将 sql 语句修改为修改 int 类型的时间戳
	DeletedAt *int `json:"deletedAt"`
}

func updateTimestampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedAt"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifiedTimeField, ok := scope.FieldByName("ModifiedAt"); ok {
			if modifiedTimeField.IsBlank {
				modifiedTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimestampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedAt", time.Now().Unix())
	}
}

// BeforeCreate will set a UUID rather than numeric ID.
func (base *Base) BeforeCreate(scope *gorm.Scope) error {
	u1 := uuid.NewV4()
	scope.SetColumn("ID", u1)
	// scope.SetColumn("CreatedAt", time.Now().Unix())
	updateTimestampForCreateCallback(scope)

	return nil
}

// BeforeUpdate gorm的Callbacks，可以将回调方法定义为模型结构的指针，在创建、更新、查询、删除时将被调用，如果任何回调返回错误，gorm 将停止未来操作并回滚所有更改。
func (base *Base) BeforeUpdate(scope *gorm.Scope) error {
	// scope.SetColumn("ModifiedAt", time.Now().Unix())
	updateTimestampForUpdateCallback(scope)

	return nil
}

// BeforeDelete ...
// func (base *Base) BeforeDelete(scope *gorm.Scope) error {
// 	if !scope.HasError() {
// 		nowTime := time.Now().Unix()
// 		if deleteTimeField, ok := scope.FieldByName("DeletedAt"); ok {
// 			if deleteTimeField.IsBlank {
// 				// FIXME: 无法设置时间戳，仍然会由 gorm 自行设置成字符串类型的时间，导致软删除失败
// 				deleteTimeField.Set(nowTime)
// 			}
// 		}
// 	}
// 	return nil
// }

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

// 参考源码 https://github.com/jinzhu/gorm/blob/master/callback_delete.go
// deleteCallback used to delete data from database or set deleted_at to current time (when using with soft delete)
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}

		deletedAtField, hasDeletedAtField := scope.FieldByName("DeletedAt")

		if !scope.Search.Unscoped && hasDeletedAtField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedAtField.DBName),
				// 设置时间戳
				scope.AddToVars(time.Now().Unix()),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
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

	// 使用新函数`deleteCallback`替换回调`gorm:delete`用于删除过程
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

// CloseDB 关闭数据库连接
func CloseDB() {
	defer db.Close()
}
