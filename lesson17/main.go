package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()

	//路由引擎
	router := gin.Default()

	//RESTful风格
	router.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"GET",
		})
	})

	router.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"POST",
		})
	})

	router.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"PUT",
		})
	})

	router.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"method":"DELETE",
		})
	})

	//Any:请求方法的大集合/大杂烩
	router.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case http.MethodGet:
			c.JSON(http.StatusOK,gin.H{
				"method":"GET",
			})
		case http.MethodPost:
			c.JSON(http.StatusOK,gin.H{
				"method":"POST",
			})
		case http.MethodPut:
			c.JSON(http.StatusOK,gin.H{
				"method":"PUT",
			})
		case http.MethodDelete:
			c.JSON(http.StatusOK,gin.H{
				"method":"DELETE",
			})
		default:
			c.JSON(http.StatusNotFound,gin.H{
				"message":"后端未处理当前请求方式",
			})
		}
	})

	//404
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound,gin.H{
			"message":"当前路由未映射",
		})
	})

	//路由组 多用于区分不同的业务线或API版本
	//把公用的前缀提取出来，创建一个路由组
	vidoGroup:=router.Group("/video")
	{
		vidoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"video/index"})
		})
		vidoGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"video/xx"})
		})
		vidoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"video/oo"})
		})
	}

	//商城
	shopGroup := router.Group("/shop")
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"/shop/index"})
		})
		shopGroup.GET("/xx", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"/shop/xx"})
		})
		shopGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{"msg":"/shop/oo"})
		})
	}

	router.Run(":9000")
}
