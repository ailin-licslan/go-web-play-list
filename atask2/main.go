package main

import (
	"fmt"
	"sync"
)

func main() {

	//TEST Exercise 1  指针
	b := 20
	a := test(&b)
	fmt.Println(a) //30

	c := 30
	d := 40
	swap(c, d)     //没有交换
	fmt.Println(c) //30
	fmt.Println(d) //40

	e := 50
	f := 60
	swap2(&e, &f)  //交互  传指针指向的地址
	fmt.Println(e) //60
	fmt.Println(f) //50

	arr := []int{1, 2, 3}
	fmt.Println(doubleValue(&arr)) // [2 4 6]

	//TEST Exercise 2  goroutine
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		printEven()
	}()
	go func() {
		defer wg.Done()
		printOdd()
	}()
	wg.Wait()

	//TEST Exercise 3 面向对象 (1.计算面积   2.员工信息)
	// 创建矩形实例
	rectangle := Rectangle{Width: 5, Height: 3}
	// 调用矩形的面积和周长方法
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectangle.Perimeter())
	// 创建圆形实例
	circle := Circle{Radius: 4}
	// 调用圆形的面积和周长方法
	fmt.Printf("Circle Area: %.2f\n", circle.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.Perimeter())

	//创建一个Employee实例
	employee := Employee{
		EmployeeID: 1,
		Person: Person{
			Name: "LIN",
			Age:  20,
		},
	}
	employee.printInfo()

	//TEST Exercise 4 channel  TODO

	//TEST Exercise 3 锁机制  TODO

}
