package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	r.GET("user/search/:username/:address", func(c *gin.Context) {
		// 这两个都不能忽略，不然就访问不到该路由
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	r.Run()
}
