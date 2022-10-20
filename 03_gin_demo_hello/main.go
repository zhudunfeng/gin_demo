package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayhello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Golang!",
	})
}

func main() {
	gin.ForceConsoleColor()
	r := gin.Default() //返回默认的路由引擎

	//指定用户使用GET请求访问/hello时,执行sayhello这个函数
	r.GET("/hello", sayhello)

	r.GET("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	//启动服务
	r.Run(":9090")
}
