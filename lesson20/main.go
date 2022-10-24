package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"` //自定义列名
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"` // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	Num          int     `gorm:"AUTO_INCREMENT"` // 设置 num 为自增类型
	Address      string  `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"` // 忽略本字段
}

//自定义表名
//func (u *User) TableName() string {
//	return  "my_user"
//}


func main() {
	//GORM还支持更改默认表名称规则
	//可以增加统一前缀或后缀
	gorm.DefaultTableNameHandler = func (db *gorm.DB, defaultTableName string) string  {
		return "prefix_" + defaultTableName;
	}

	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// 禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	//db.SingularTable(true)

	//自动迁移
	db.AutoMigrate(&User{})

	// 使用User结构体创建名为`test_users`的表
	db.Table("test_users").CreateTable(&User{})
}
