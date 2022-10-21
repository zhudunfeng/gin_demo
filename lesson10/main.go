package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/json", func(c *gin.Context) {
		//方法一：使用map
		//data:=map[string]interface{}{
		//	"name":"孙悟空",
		//	"message":"hello world!",
		//	"age":18,
		//}

		data:=gin.H{
			"name":    "孙悟空",
			"message": "hello world!",
			"age":     18,
		}
		c.JSON(http.StatusOK,data)
	})

	//方法2：结构体
	type msg struct {
		Name string `json:"name"`
		Message string	`json:"message"`
		Age int	`json:"age"`
	}

	router.GET("/struct", func(c *gin.Context) {
		data:= msg{
			Name:"张三",
			Message:"Hello",
			Age:18,
		}
		c.JSON(http.StatusOK,data)
	})


	err := router.Run(":9000")
	if err != nil {
		fmt.Println("Http Server error:",err)
		return
	}
}
