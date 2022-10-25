package main

import (
	"github.com/gin-gonic/gin"
	"go_front/dao"
	"go_front/model"
	"go_front/routers"
)

func main() {
	gin.ForceConsoleColor()

	//创建数据库
	//sql: create database bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.Close()
	//模型绑定
	dao.DB.AutoMigrate(&model.Todo{})
	router := routers.SetupRouter()

	router.Run(":9000")
}
