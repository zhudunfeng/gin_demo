package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()

	router.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently,"http://www.sogou.com")
	})

	router.GET("/a", func(c *gin.Context) {
		c.Request.URL.Path = "/b"
		router.HandleContext(c)
	})

	router.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"b",
		})
	})

	router.Run(":9000")
}
