package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//首先定义模型：
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// user1 := User{Name: "qimi", Age: 20, Active: true}
	// user2 := User{Name: "uzi", Age: 18, Active: false}

	// db.Create(&user2)
	// db.Create(&user1)

	var user User
	db.First(&user)

	//6.更新
	// user.Name = "1"
	// db.Save(&user) //默认修改所有字段,其他没设的字段就没了

	// db.Debug().Delete(&user)
	// db.Unscoped().Delete(&user)
}
