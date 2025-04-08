package main

import "fmt"

//func swap(a int, b int) {
//	var temp int
//	temp = a
//	a = b
//	b = temp
//}

// 传递指针
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa //temp = main::a
	*pa = *pb  // main::a = main::b
	*pb = temp // main::b = temp
}

func main() {
	var a int = 10
	var b int = 20

	// 传 &a &b 地址
	swap(&a, &b)
	// 传的值 不影响外面的变量
	//swap(a, b)

	fmt.Println("a = ", a, " b = ", b)

	var p *int

	p = &a

	fmt.Println(&a)
	fmt.Println(p)

	var pp **int //二级指针

	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)
}
