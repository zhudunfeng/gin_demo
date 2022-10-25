package routers

import (
	"github.com/gin-gonic/gin"
	"go_front/controller"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()
	//放行静态资源
	router.Static("/static", "./static")

	//定义模板
	//解析模板
	router.LoadHTMLGlob("templates/*")
	//渲染模板
	router.GET("/", controller.IndexHandler)

	//v1
	v1Group := router.Group("/v1")
	{
		//待办事项
		//添加
		v1Group.POST("/todo", controller.CreateOneTodo)
		//查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//查看某一个待办事项
		v1Group.GET("/todo/:id", controller.GetTodoById)
		//修改某一个待办事项
		v1Group.PUT("/todo/:id", controller.UpdateTodoById)
		//删除某一个待办事项
		v1Group.DELETE("/todo/:id", controller.DeleteTodoById)
	}
	return router
}
