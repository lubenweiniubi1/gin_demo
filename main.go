package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		//方法一:使用map
		// data := map[string]interface{}{
		// 	"name":    "小王子",
		// 	"message": "hello world",
		// 	"age":     18,
		// }
		//gin.H就是 map[string] interface {} 这个interface{}就是空接口类型，可以存储任何值
		c.JSON(http.StatusOK, gin.H{
			"name":    "小王子",
			"message": "hello world",
			"age":     18,
		})
	})
	r.Run()
}
