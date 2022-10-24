package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name *string `gorm:"default:"阿敦"` //*string 允许空值
	Age sql.NullInt64
	//Age int64
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//延迟回收资源
	defer db.Close()

	//2.把模型与数据库的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	//u := User{Name: new(string), Age: sql.NullInt16{0,true}} //在代码层面创建一个User对象
	name := "aaaaa"
	u := User{Name: &name,Age: sql.NullInt64{20,true}} //在代码层面创建一个User对象
	fmt.Println(db.NewRecord(&u))  //判断逐渐是否为空 true
	db.Debug().Create(&u)          //在数据库中创建一条 adun 18 的记录
	fmt.Println(db.NewRecord(&u))  //判断主键是否为空 false

}
