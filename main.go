package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		//获取浏览器query 参数
		username := c.DefaultQuery("username", "小丸子") //带默认值,娶不到就是默认值
		address := c.Query("address")

		name, ok := c.GetQuery("name") //娶不到就是sb
		if !ok {
			name = "sb"
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
			"name":     name,
		})
	})

	r.Run()
}
