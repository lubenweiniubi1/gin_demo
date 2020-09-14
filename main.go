package main

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//定义模型
type User struct {
	gorm.Model   //内嵌gorm.Model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号（member number )唯一 不为空
	Num          int     `gorm:"AUTO_INCREMENT`   //设置num为自增类型
	Address      string  `gorm:"index:addr"`      //给address字段创建名为addr的索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

type Animal struct {
	AnimalID int64  `gorm:"primary_key"`
	Name     string `gorm:"column:best_go"`
	Age      int64
}

func (Animal) TableName() string {
	return "qimi"
} //设置自定义表明

func main() {
	//连接数据库
	db, err := gorm.Open("mysql", "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Gorm还支持更改默认表名称规则
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName //自定义的表明就不会加
	}

	db.SingularTable(true) //禁用表的复数

	db.AutoMigrate(&User{}) //创建表
	db.AutoMigrate(&Animal{})

	db.Table("uziniupi").CreateTable(&User{}) //使用user结构体创建一个名叫"uziniupi"的表

}
