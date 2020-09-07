package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//handlefunc 类型
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//定义一个中间件 m1
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
}

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	// func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	r.GET("/index", m1, indexHandler)

	r.Run(":8080")
}
