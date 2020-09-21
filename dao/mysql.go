package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4&loc=Local&parseTime=true"
	DB, err = gorm.Open("mysql", dsn) //不要冒号等于重复生命，已经声明全局了
	if err != nil {
		return
	}

	return DB.DB().Ping() //ping的通就返回nil
}
