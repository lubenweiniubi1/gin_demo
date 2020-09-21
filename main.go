package main

import (
	"gin_demo/controller"
	"gin_demo/dao"
	"gin_demo/models"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
) //这里涉及到gorm 版本变更问题，目前使用视频里的东西，

//Todo Model

func main() {
	//创建数据库
	//创建bubble
	//连接数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	//模型绑定
	dao.DB.AutoMigrate(&models.Todo{})
	r := gin.Default()
	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", controller.AddTodo)
		//查看所有代办事项
		v1Group.GET("/todo", controller.GetTodoList)
		//修改
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		//删除
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}

	r.Run(":9000")

}
