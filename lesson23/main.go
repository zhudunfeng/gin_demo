package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type User struct {
	gorm.Model //ID CreateAt UpdateAt DeleteAt
	Name string
	Age int64
	Active bool
}

func main() {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//自动迁移
	db.AutoMigrate(&User{})

	//插入数据
	//db.Create(&User{Name:"阿敦3",Age: 18,Active: true})
	//db.Create(&User{Name:"阿敦4",Age: 20,Active: true})

	//查询
	var user User
	db.First(&user)
	fmt.Printf("%#v\n",user)

	//更新
	user.Name = "永远的神"
	user.Age = 99
	db.Debug().Save(&user) //默认会修改所有的字段

	db.Debug().Model(&user).Update("name","ADUN")

	m1:=map[string]interface{}{
		"name":"墩墩",
		"age":18,
		"active":false,
	}
	db.Debug().Model(&user).Updates(m1) //m1列出的所有字段都会更新
	db.Debug().Model(&user).Select("age").Updates(m1) //只更新age字段
	db.Debug().Model(&user).Omit("active").Updates(m1) //排除m1中的active更新其他的字段


	//更新条件
	user1:=User{
		Name:"ADUN@@@",
		Active: false, //这个不能更新到数据库
	}
	// 使用 struct 更新时，只会更新非零值字段，若想更新所有字段，请使用map[string]interface{}
	rowCount := db.Debug().Model(&User{}).Where("age=?", 20).Updates(user1).RowsAffected
	fmt.Println(rowCount)

	rowCount = db.Debug().Model(&User{}).Where("age=?", 20).Updates(map[string]interface{}{
		"name":"ADUN&&&",
		"active": false,
	}).RowsAffected
	fmt.Println(rowCount)

	//使用SQL表达式更新
	rowCount = db.Debug().Model(&User{}).Where("age=?", 20).
		UpdateColumn("age", gorm.Expr("age+?", 1)).RowsAffected
	fmt.Println(rowCount)

	fmt.Println("11111")
	fmt.Println("222222")
}
