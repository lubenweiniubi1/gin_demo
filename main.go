package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	r.POST("/login", func(c *gin.Context) {
		// username := c.PostForm("username")
		password := c.PostForm("password")
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username, ok := c.GetPostForm("username")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"ok":       ok,
		})
	})

	r.Run()
}
