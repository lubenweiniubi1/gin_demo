package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
) //这里涉及到gorm 版本变更问题，目前使用视频里的东西，

var (
	DB *gorm.DB
)

//Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMySQL() (err error) {
	dsn := "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4&loc=Local&parseTime=true"
	DB, err = gorm.Open("mysql", dsn) //不要冒号等于重复生命，已经声明全局了
	if err != nil {
		return
	}

	return DB.DB().Ping() //ping的通就返回nil
}

func main() {
	//创建数据库
	//创建bubble
	//连接数据库
	err := initMySQL()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//模型绑定
	DB.AutoMigrate(&Todo{})
	r := gin.Default()
	//v1
	v1Group := r.Group("v1")
	{
		//代办事项
		//添加
		v1Group.POST("/todo", func(c *gin.Context) {
			//前端页面填写代办事项，会发请求到这里
			//1.把请求搞出来
			var todo Todo
			c.BindJSON(&todo)
			//2.存入数据库
			if err = DB.Create(&todo).Error; err != nil {
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

		})
		//查看所有代办事项
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		//查看某一个待办事项
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		//删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	r.Run()

}
