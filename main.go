package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//首先定义模型：
type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.AutoMigrate(&User{})

	// user := User{Name: "qimi", Age: 418}
	// user1 := User{Name: "uzi", Age: 18}

	// db.Create(&user)
	// db.Create(&user1)

	if false {
		//一般查询
		var user User
		db.First(&user)
		// fmt.Println(user)

		var users []User
		db.Debug().Find(&users) //查询所有记录
		// fmt.Printf("users:%#v\n", users)

		//where
		var user2 User
		db.Where("name = ?", "qimi").First(&user2)
		fmt.Printf("users:%#v\n", user2)

		// db.Where("name = ?", "qimi").First(&user)
		// fmt.Printf("users:%#v\n", user) //好像是多携程的，别用同一个变凉了
	}

	//FirstOrInit
	var user1 User
	db.Attrs(User{Age: 100}).FirstOrInit(&user1, User{Name: "non_existing"})
	// fmt.Printf("user:%#v\n", user1) //Name:"non_existing", Age:100

	var user3 User
	// 未找到,Assgin没卵用
	db.Assign(User{Age: 20}).FirstOrInit(&user3, User{Name: "uzi"})
	fmt.Printf("user:%#v\n", user3) //Name:"non_existing", Age:100

}
