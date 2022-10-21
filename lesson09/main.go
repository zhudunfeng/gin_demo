package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func main() {
	gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	//处理静态文件 css js 图片s
	router.Static("/xxx","./statics")
	//自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"safe": func(str string)template.HTML {
			return template.HTML(str)
		},
	})

	//定义模板
	//解析模板
	//router.LoadHTMLFiles("templates/index.tmpl","./template/posts/index.tmpl")
	//router.LoadHTMLGlob("templates/**/*.tmpl")
	router.LoadHTMLGlob("templates/**/*")
	//渲染模板
	router.GET("/index", func(c *gin.Context) {
		//HTTP请求
		c.HTML(http.StatusOK,"index.tmpl",gin.H{
			"title":"zhudunfeng.github.io",
		})
	})
	router.GET("posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"posts/index.tmpl",gin.H{
			"title":"posts/index.tmpl",
		})
	})
	router.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK,"users/index.tmpl",gin.H{
			"title":"<a href='https://zhudunfeng.github.io/'>ADUN</a>",
		})
	})

	// 返回从网上下载的模板
	router.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK,"home.html",nil)
	})
	router.Run(":9000") //启动server
}
