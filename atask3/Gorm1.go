package main

/**
题目1：模型定义
假设你要开发一个博客系统，有以下几个实体： User （用户）、 Post （文章）、 Comment （评论）。
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章），
Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
//首先我手动设计了下 看看自动生成的和手动的差异

// User 模型定义
//type User struct {
//	gorm.Model //一些表的常见base信息
//	Username   string
//	Phone      string
//	Posts      []Post `gorm:"foreignKey:UserID"` //一对多关系
//}
//
//// Post 模型定义
//type Post struct {
//	gorm.Model //一些表的常见base信息
//	Title      string
//	Content    string
//	UserID     uint
//	User       User      `gorm:"references:ID"`
//	Comments   []Comment `gorm:"foreignKey:PostID"` //一对多关系
//}
//
//// Comment 模型定义
//type Comment struct {
//	gorm.Model
//	Content string
//	PostID  uint
//	UserID  uint
//	Post    Post `gorm:"references:ID"`
//	User    User `gorm:"references:ID"`
//}

//func main() {
//	// 替换为你的数据库连接信息
//	dsn := "root:123456@tcp(192.168.0.155:3306)/blog_auto?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	// 自动迁移模式，创建表
//	_ = db.AutoMigrate(&User{}, &Post{}, &Comment{})
//}
