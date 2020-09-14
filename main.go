package main

import (
	"database/sql"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//首先定义模型：
type User struct {
	ID int64
	// Name *string `gorm:"default:'小王八'"` //使用指针方式实现零值存入数据库
	Name sql.NullString `gorm:"default:'小王八'"` //使用Scanner/Valuer接口方式实现零值存入数据库
	Age  int64
} //修改 但是表结构没改过来,就删掉吧

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2. 把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	// user := User{Age: 29}           //在代码层创建一个User对象,不写名字 默认就是空串
	// user := User{Name: "", Age: 52} //这里的名字还是小王八 ,用的默认值,也就是说你输入零值 和不写没有区别
	// user := User{Name: new(string), Age: 52} //当我不传的时候 你使用小王八 ,我传了值的时候 我传什么你写什么 ,new(string) 获取空字符串指针
	user := User{Name: sql.NullString{String: "", Valid: true}, Age: 512412} //...不可能用这个了
	fmt.Println(db.NewRecord(user))                                          //检查主键是否为空，主键一般由数据库来生成，如果为true说明还没被创建，如果为false，说明数据库里由这个东西了，不能再创建了
	db.Debug().Create(&user)                                                 //在数据库中创建了一条 七米 78的记录
	fmt.Println(db.NewRecord(user))

}
