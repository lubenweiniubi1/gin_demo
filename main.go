package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//_ 表示没有直接用到他

//UserInfo --> 数据表
type UserInfo struct {
	ID    uint
	Name  string
	Gener string
	Hobby string
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//创建表  自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	//创建数据行
	u1 := UserInfo{1, "七米", "男", "蛙泳"} //生成的数据表的名字为user_infos
	u2 := UserInfo{2, "无敌", "女", "足球"}
	db.Create(&u1) //结构体对象比较大就可以传指针进来
	db.Create(&u2)

	//查询
	// var u UserInfo
	var u = new(UserInfo)
	db.First(&u) //会修改变量 u,来保存查询的值
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	//更新
	db.Model(&u).Update("hobby", "双色球")
	//删除
	db.Delete(&u)
}
