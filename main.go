package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//handlefunc 类型
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	name, _ := c.Get("name")
	c.JSON(http.StatusOK, gin.H{
		"msg": name,
	})
}

//定义一个中间件 m1
//统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	go funcXX(c.Copy()) //在funcXX中只能使用c的拷贝,使用只读的对象
	c.Next()            //调用后续的处理函数,这里十分重要！！！
	// c.Abort()//阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Set("name", "qimi") //将数据放到c的上下文中给后面的使用
	c.Next()
	// c.Abort() //组织后续的处理函数，这里直接打印m2 out
	fmt.Println("m2 out ...")
}

// func authMiddleware(c *gin.Context) {
// 	//是否登陆判断
// 	//if 是登陆用户
// 	//c.Next
// 	//else
// 	//c.Abort
// }

//通常会用闭包 ，可以自己加参数
func authMiddleware(doCheck bool) gin.HandlerFunc {
	//连接数据库
	//或者一些其他的准备工作
	return func(c *gin.Context) {
		if doCheck {
			//存放具体的逻辑
			//是否登陆判断
			//if 是登陆用户
			//c.Next
			//else
			//c.Abort
		} else {
			c.Next()
		}

	}
}

func StatCost(part string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("你吗死了")
		c.Next()
		fmt.Println("我是你爹")
	}
}

//go mod tidy 看依赖包是否在go mod 中列出来
func main() {
	// r := gin.Default()                  //默认使用了Logger()和Recovery()中间件
	r := gin.New()                      //不想使用两个东西的话
	r.Use(m1, m2, authMiddleware(true)) //全局注册中间件函数 ,打印顺序：m1 in ... ，m2 in ... ， index ,m2 out ...,cost:996.8µs  ,m1 out ...
	// func (group *RouterGroup) GET(relativePath string, handlers ...HandlerFunc) IRoutes {
	// r.GET("/index", m1, indexHandler)
	r.GET("/index", indexHandler)
	r.GET("/index1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"messag": "这里是index2",
		})
	})

	//为路由组注册
	shopGroup := r.Group("/xx", StatCost("ning "))
	{
		shopGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"msg": "这里是路由组",
			})
		})
	}

	r.Run(":8080")
}

func funcXX(c *gin.Context) {
	fmt.Println("因为在这里如果直接使用c，会导致后面的处理程序的值是不确定的，并发不安全")
}
