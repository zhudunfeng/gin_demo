package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	//GET请求 URL ？ 后面的是querystring参数
	//key=value ，多个key=value用&连接
	//eq: /web/query=阿敦&age=18

	router := gin.Default()

	router.GET("/web", func(c *gin.Context) {
		//获取浏览器那边发请求携带的querystring参数
		//name:=c.Query("name") //通过Query获取请求中携带的querystring参数
		age := c.Query("age")
		//name:=c.DefaultQuery("name","somebody") //取不到就用指定的默认值
		name, ok := c.GetQuery("name") //取到返回（值,true）,渠道到返回（"",false）
		if !ok {
			//取不到
			name="somebody"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	err := router.Run(":9000")
	if err != nil {
		fmt.Println("Http Server start failed:", err)
		return
	}
}
