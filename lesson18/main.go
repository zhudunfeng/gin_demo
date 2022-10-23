package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

//定义一个耗时组件
func StatCost() gin.HandlerFunc {
	//连接数据库
	//或者其他一些准备工作
	return func(c *gin.Context) {
		fmt.Println("Stat in....")
		//计时
		start := time.Now()

		// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		c.Set("name", "阿敦")

		//当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
		go func() {
			//c1 := c.Copy()
			//time.Sleep(time.Second*2)
			fmt.Println("OKK")
		}()

		c.Next() //调用后续的处理函数
		//c.Abort() //阻止调用后续的处理函数
		cost := time.Since(start)
		fmt.Printf("Cost time:%v\n", cost)

		fmt.Println("Stat in....")
	}
}

//权限认证
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者其他一些准备工作
	return func(c *gin.Context) {
		fmt.Println("authMiddleware in...")
		//存放具体的逻辑
		//是否登录的判断
		//if 是登录用户
		//c.Next()
		//else
		//c.Abort()

		name := c.Query("name")
		if "adun" == name {
			c.Next()
		} else {
			c.Abort()
		}

		fmt.Println("authMiddleware out...")
	}
}

func main() {
	gin.ForceConsoleColor()

	router := gin.Default() //默认使用了Logger和Recovery中间件
	//router:=gin.New()//如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由

	//全局路由注册组件
	router.Use(StatCost(), authMiddleware(true))

	//为单个路由注册组件
	//router.GET("/index",StatCost(), indexHandler)
	router.GET("/index", indexHandler)
	router.GET("/index2", func(c *gin.Context) {
		name, ok := c.Get("name")
		if !ok {
			fmt.Println("name 不存在")
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	router.GET("/index3")

	err := router.Run(":9000") //默认8080
	if err != nil {
		fmt.Println("Http Server start failed:", err)
		return
	}
}
