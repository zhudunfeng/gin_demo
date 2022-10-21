package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserInfo struct {
	Name string `uri:"name" form:"name" json:"name" binding:"required"`
	Age  int	`uri:"age" form:"age" json:"age"`
}

//演示参数绑定
func main() {
	gin.ForceConsoleColor()
	router := gin.Default()

	//定义模板
	//解析模板
	router.LoadHTMLFiles("./index.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"message": "ok",
		})
	})

	router.GET("/user1", func(c *gin.Context) {
		username := c.Query("username")
		age := c.Query("age")
		tmp, err := strconv.ParseInt(age, 10, 64)

		if err != nil {
			fmt.Println("strconv.ParseInt error:", err)
			return
		}
		u := UserInfo{
			Name: username,
			Age:  int(tmp),
		}
		c.JSON(http.StatusOK,u)
	})

	//querystring参数绑定
	router.GET("/user2", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的实例u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%v\n",u)
			c.JSON(http.StatusOK,u)
		}

	})

	//表单参数绑定
	router.POST("/user2", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的实例u
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%v\n",u)
			c.JSON(http.StatusOK,u)
		}
	})

	//param参数绑定
	router.GET("/user3/:name/:age", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的实例u
		err := c.ShouldBindUri(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%v\n",u)
			c.JSON(http.StatusOK,u)
		}

	})

	//json绑定
	router.POST("/user4", func(c *gin.Context) {
		var u UserInfo //声明一个UserInfo类型的实例u
		err := c.ShouldBindJSON(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}else{
			fmt.Printf("%v\n",u)
			c.JSON(http.StatusOK,u)
		}

	})

	router.Run(":9000")
}
