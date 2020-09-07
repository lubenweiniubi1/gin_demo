package main

import (
	"fmt"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	r := gin.Default()
	r.LoadHTMLFiles("./index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	r.POST("/upload", func(c *gin.Context) {
		//从请求中读取文件
		fileReader, err := c.FormFile("f1") //从请求中获取携带的参数
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"err": err.Error(),
			})
		} else {
			//将文件保存到本地
			// dst := fmt.Sprintf("./%s",f.Filename)
			dst := path.Join("./", fileReader.Filename)
			fmt.Printf(fileReader.Filename)         //自拍.jpg
			_ = c.SaveUploadedFile(fileReader, dst) //短线就不加冒号了 ，这里不处理错误
			c.JSON(http.StatusOK, gin.H{
				"status": "OK",
			})
		}
		//将读取出的文件保存到本地
	})
	r.Run(":8080")
}
