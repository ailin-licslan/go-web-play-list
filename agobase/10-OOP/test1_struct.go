package main

import "fmt"

// 声明一种行的数据类型 myint， 是int的一个别名
type myint int

// 定义一个结构体
type BookV1 struct {
	title string
	auth  string
}

func changeBook(book BookV1) {
	//传递一个book的副本
	book.auth = "666"
}

func changeBook2(book *BookV1) {
	//指针传递
	book.auth = "777"
}

func main() {
	/*
		var a myint = 10
		fmt.Println("a = ", a)
		fmt.Printf("type of a = %T\n", a)
	*/

	var book1 BookV1
	book1.title = "Golang"
	book1.auth = "zhang3"

	fmt.Printf("%v\n", book1)

	changeBook(book1)

	fmt.Printf("%v\n", book1)

	changeBook2(&book1)

	fmt.Printf("%v\n", book1)
}
