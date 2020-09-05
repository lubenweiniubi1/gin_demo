package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
} //应用反射，让别人拿到你的字段 ，首字母大写 ，并指定tag ,requred表示必须的，不传会报错 ; json字段表示 如果穿的是json文件，就根据这个来选择 username属性，

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		// 这两个都不能忽略，不然就访问不到该路由

		var user UserInfo //生命一个user【info类型变量

		err := c.ShouldBind(&user) // 函数都是值拷贝，这里传的是个拷贝，应该传一个地址
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Println(user)
	})
	r.POST("/user", func(c *gin.Context) {
		var user UserInfo

		err := c.ShouldBind(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
		fmt.Println(user)

	})

	r.Run()
}
