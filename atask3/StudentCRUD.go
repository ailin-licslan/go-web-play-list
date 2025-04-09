package main

import (
	"fmt"
	"github.com/jinzhu/gorm"                  //"gorm.io/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //old gorm jin zhu team  (老版本的)   建议使用新的版本"gorm.io/driver/mysql"
)

/**
题目1：基本CRUD操作
假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
要求 ：
编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。
*/

type Student struct {
	ID     uint `gorm:"primary_key;AUTO_INCREMENT"`
	Name   string
	Age    int
	Gender string
}

func main() {
	db, err := gorm.Open("mysql",
		"root:123456@(192.168.0.155:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// 自动迁移
	db.AutoMigrate(&Student{})
	u0 := Student{Name: "张三", Age: 20, Gender: "三年级"}
	u1 := Student{Name: "LIN", Age: 20, Gender: "9"}
	u2 := Student{Name: "HI", Age: 18, Gender: "5"}
	u3 := Student{Name: "H", Age: 13, Gender: "3"}

	// 创建记录
	tx := db.Begin()
	if tx.Error != nil {
		panic("开启事务失败")
	}
	// 执行多个数据库操作
	db.Create(&u0)
	db.Create(&u1)
	db.Create(&u2)
	db.Create(&u3)
	// 提交事务
	tx.Commit()

	// 查询
	var u = new(Student)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var students []Student
	db.Where("age > 18").Find(&students)
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d, Gender: %s\n", student.ID, student.Name, student.Age, student.Gender)
	}

	var uu Student
	db.Find(&uu, "name=?", "LIN")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&Student{}).Where("name =?", "LIN").Update("age", 21)
	//db.Model(&u).Update("name", "Jack")

	// 删除
	db.Where("age < 15").Delete(&Student{})
	//db.Delete(&u)
}
