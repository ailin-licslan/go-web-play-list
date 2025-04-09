package main

//import (
//	"fmt"
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//)
//
///**
//题目2：关联查询
//基于上述博客系统的模型定义。
//要求 ：
//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
//编写Go代码，使用Gorm查询评论数量最多的文章信息。
//*/
//
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
//
//// 查询某个用户发布的所有文章及其对应的评论信息
//func queryUserPostsAndComments(db *gorm.DB, userID uint) {
//	var user User
//	err := db.Preload("Posts.Comments").Where("id =?", userID).First(&user).Error
//	if err != nil {
//		fmt.Printf("查询用户时出错: %v", err)
//		return
//	}
//
//	fmt.Printf("用户 %s 的文章及评论信息:", user.Username)
//	for _, post := range user.Posts {
//		fmt.Printf("文章标题: %s", post.Title)
//		for _, comment := range post.Comments {
//			fmt.Printf("  评论内容: %s", comment.Content)
//		}
//	}
//}
//
//// 查询评论数量最多的文章信息
//func queryMostCommentedPost(db *gorm.DB) {
//	var post Post
//	subQuery := db.Model(&Comment{}).Select("post_id, count(*) as comment_count").Group("post_id").Order("comment_count desc").Limit(1)
//	err := db.Model(&Post{}).Joins("JOIN (?) as sub ON posts.id = sub.post_id", subQuery).First(&post).Error
//	if err != nil {
//		fmt.Printf("查询评论最多的文章时出错: %v", err)
//		return
//	}
//
//	fmt.Printf("评论数量最多的文章标题: %s", post.Title)
//}
//
//func main() {
//	// 替换为你的数据库连接信息
//	dsn := "root:123456@tcp(192.168.0.155:3306)/blog_auto?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		panic("failed to connect database")
//	}
//
//	// 查询某个用户（假设用户ID为1，可自行替换）发布的所有文章及其对应的评论信息
//	queryUserPostsAndComments(db, 1)
//
//	// 查询评论数量最多的文章信息
//	queryMostCommentedPost(db)
//}
