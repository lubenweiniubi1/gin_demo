package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()

	// r.HEAD()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "GET",
		})
	})
	r.POST("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})
	r.DELETE("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.PUT("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.Any("/anny", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code": "404",
		})
	})

	//我们可以将拥有共同URL前缀的路由划分为一个路由组。习惯性一对{}包裹同组的路由，这只是为了看着清晰，你用不用{}包裹功能上没什么区别。
	//视频的首页和详情页
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})

		videoGroup.GET("/cc", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})

		videoGroup.GET("/dd", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})
		videoGroup.GET("/oo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "/video/index",
			})
		})

		// 嵌套路由组
		xx := videoGroup.Group("xx")
		xx.GET("/oo", func(c *gin.Context) {})
	}

	r.GET("/shop/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "/shop/index",
		})
	})

	//any 请求大杂烩
	r.Any("/user", func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			return
		case http.MethodPost: //尽量用常量
			return
		}
	})
	r.Run(":8080")
}
