package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //返回默认路由引擎

	//指定用户使用GET请求访问/ping 时，执行func。。这个函数
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{ //H是格 map，自己点进去看
			"method": "get",
		})
	})
	r.POST("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.PUT("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	//启动服务
	r.Run()

}
