##  GORM

#### 什么是ORM？

Object Relational Mapping（对象  关系 映射）

+ Object:程序中的对象/实例，比如go中的结构体实例
+ Relational:关系型数据库，例如mysql
+ Mapping：一个实例对应数据库中的一条数据

```go
type UserInfo struct {
    ID uint 
    Name string 
    Gender string
}

func main(){
    u1:=UserInfo{1,"七米","男"}
    //将u1数据存入数据库
   // insert into userinfo values(1,"七米","男") sql语句
    orm.Create(&u1) //orm工具提供的orm语句
}
```

![go1](D:\gin_demo\go1.png)

> https://www.liwenzhou.com/posts/Go/gorm/