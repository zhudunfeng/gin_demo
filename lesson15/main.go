package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
)

func main() {

	gin.ForceConsoleColor()

	router := gin.Default()

	//处理静态资源
	router.Static("/static","./statics")

	router.LoadHTMLFiles("./index.html", "./multiFileUpload.html")

	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.GET("/files", func(c *gin.Context) {
		c.HTML(http.StatusOK, "multiFileUpload.html", nil)
	})

	// 处理multipart forms提交文件时默认的内存限制是32 MiB
	// 可以通过下面的方式修改
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	//上传单个文件
	router.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保存到本地（服务端本地）
			//dst:=fmt.Sprintf("./%s",file.Filename)
			dst := path.Join("./", file.Filename)
			//保存文件
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})

	//上传多个文件
	router.POST("/uploadFiles", func(c *gin.Context) {
		//从请求中读取文件
		// Multipart form
		form, err := c.MultipartForm()
		files := form.File["files"]

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			//将读取到的文件保存到本地（服务端本地）
			//保存文件
			for _,file:=range files{
				//将读取到的文件保存到本地（服务端本地）
				//dst:=fmt.Sprintf("./%s",file.Filename)
				dst := path.Join("./", file.Filename)
				//保存文件
				c.SaveUploadedFile(file, dst)
			}
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
	})

	router.Run(":9000")
}
