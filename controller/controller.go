package controller

import (
	"fmt"
	"gin_demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
  url -> controller --> logic --> model
  请求来了 -> 控制器 -->业务逻辑 --->模型层增删改查
*/
func AddTodo(c *gin.Context) {
	//前端页面填写代办事项，会发请求到这里
	//1.把请求搞出来
	var todo models.Todo
	c.BindJSON(&todo)
	//2.存入数据库
	err := models.CreateATodo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 2000,
			"msg":  "success",
			"data": todo,
		})
	}
	//3. 返回响应

}

func GetTodoList(c *gin.Context) {
	//查询todo这个表所有数据
	todoList, err := models.GetTodos()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

func UpdateATodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效的id",
		})
	}
	todo, err := models.GetATodo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return //不然会继续往下面执行
	}
	c.BindJSON(todo)
	fmt.Printf("user:%#v", &todo)
	if err = models.UpdataATodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
	}
	if err := models.DeleteATodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			id: "deleted",
		})
	}
}
