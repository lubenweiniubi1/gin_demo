 **注意：**通过tag定义字段的默认值，在创建记录时候生成的 SQL 语句会排除没有值或值为 零值 的字段。 在将记录插入到数据库后，Gorm会从数据库加载那些字段的默认值。

```go
user := User{Name: "", Age: 52} 
//打断点
db.Debug().Create(&user) 


[2020-09-14 22:41:49]  [10.98ms]  INSERT INTO `users` (`age`) VALUES (52)
```

