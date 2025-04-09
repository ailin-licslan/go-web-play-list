package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
*/

type User struct {
	gorm.Model   //一些表的常见base信息
	Username     string
	Phone        string
	Posts        []Post `gorm:"foreignKey:UserID"` //一对多关系
	ArticleCount int    `gorm:"default:0"`         //添加文章数量统计字段
}

type Post struct {
	gorm.Model    //一些表的常见base信息
	Title         string
	Content       string
	UserID        uint
	User          User      `gorm:"references:ID"`
	Comments      []Comment `gorm:"foreignKey:PostID"` //一对多关系
	CommentStatus string    `gorm:"default:有评论"`       //默认 有评论
}

// Comment 模型定义
type Comment struct {
	gorm.Model
	Content string
	PostID  uint
	UserID  uint
	Post    Post `gorm:"references:ID"`
	User    User `gorm:"references:ID"`
}

// BeforeCreate 在Post模型中添加钩子函数，在文章创建时自动更新用户的文章数量统计字段
func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	if err := tx.First(&user, p.UserID).Error; err == nil {
		user.ArticleCount++
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
	}
	return nil
}

// AfterDelete 在Comment模型中添加钩子函数，在评论删除时检查文章的评论数量，如果评论数量为0，则更新文章的评论状态为“无评论”
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var commentCount int64
	if err := tx.Model(&Comment{}).Where("post_id =?", c.PostID).Count(&commentCount).Error; err == nil {
		if commentCount == 0 {
			var post Post
			if err := tx.First(&post, c.PostID).Error; err == nil {
				post.CommentStatus = "无评论"
				if err := tx.Save(&post).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	// 替换为你的数据库连接信息
	dsn := "root:123456@tcp(192.168.0.155:3306)/blog_auto?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移模式，根据struct定义来自动创建表
	//_ = db.AutoMigrate(&User{}, &Post{}, &Comment{})

	// 这里可以进行一些测试操作，比如创建文章、删除评论等，来验证钩子函数的功能
	// 以下是简单的创建文章示例
	//newPost := Post{
	//	Title:   "测试文章",
	//	Content: "这是一篇测试文章",
	//	UserID:  1, // 假设用户ID为1
	//}
	//if err := db.Create(&newPost).Error; err != nil {
	//	log.Printf("创建文章时出错: %v", err)
	//}

	// 以下是简单的删除评论示例
	var comment Comment
	//删除第一条
	if err := db.First(&comment).Error; err == nil {
		if err := db.Delete(&comment).Error; err != nil {
			fmt.Printf("删除评论时出错: %v", err)
		}
	}
}
