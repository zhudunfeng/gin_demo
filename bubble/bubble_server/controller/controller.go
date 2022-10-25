package controller

import (
	"github.com/gin-gonic/gin"
	"go_front/model"
	service2 "go_front/service"
	"go_front/service/impl"
	"net/http"
)

var service service2.Service =&impl.ServiceImpl{}

/*
url		----> controller ---->logic 	----> model
请求来了	----> 控制器		-----> 业务逻辑	----> 模型层的增删改查
 */
func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateOneTodo(c *gin.Context) {
	//前端页面填写待办事项，点击提交，会转发请求到这里
	//1.从请求中把数据拿出来
	var todo model.Todo
	c.BindJSON(&todo)

	//2.存入数据库
	//3.返回响应
	if err := model.CreateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, &todo)
	}
}

func GetTodoList(c *gin.Context) {
	//查询todo这个表里的所有数据
	todoList, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func GetTodoById(c *gin.Context) {
	id := c.Param("id")
	//todo, err := model.GetTodoById(id)
	//s :=&service.ServiceImpl{}
	todo, err := service.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func UpdateTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	todo, err := model.GetTodoById(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	c.BindJSON(&todo)
	if model.UpdateATodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func DeleteTodoById(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效id"})
		return
	}
	if err := model.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{id: "deleted"})
	}
}
