package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	r.GET("/index", func(c *gin.Context) {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status": "ok",
		// })
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com") //这里地址都变了
	})

	r.GET("/a", func(c *gin.Context) {
		c.JSON(http.StatusPermanentRedirect, gin.H{
			"message": "你好 ，这里到路由a了",
		})
		c.Request.URL.Path = "/b" //把请求的URI修改
		r.HandleContext(c)        //继续后续处理
	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusPermanentRedirect, gin.H{
			"message": "你好 ，这里重定向到路由b了",
		})
		c.Request.URL.Path = "/c"
		r.HandleContext(c)
	})
	r.GET("/c", func(c *gin.Context) {
		c.JSON(http.StatusPermanentRedirect, gin.H{
			"message": "这里是c",
		})
	})
	r.Run(":8080")
}
