package main

import (
	"fmt"
	"math"
)

// Shape 接口定义了计算面积和周长的方法
type Shape interface {
	Area() float64      //定义一个计算面积的方法
	Perimeter() float64 //定义一个计算周长的方法
}

// Rectangle 结构体表示矩形
type Rectangle struct {
	Width  float64 //宽度
	Height float64 //高度
}

// Area 方法实现了 Shape 接口中计算矩形面积的方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 方法实现了 Shape 接口中计算矩形周长的方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 结构体表示圆形
type Circle struct {
	Radius float64 //半径
}

// Area 方法实现了 Shape 接口中计算圆形面积的方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius // Π * r^2
}

// Perimeter 方法实现了 Shape 接口中计算圆形周长的方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius //2 * Π * r
}

// Person 定义一个person结构体
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person     Person //可以写一个如果相同
	EmployeeID int
}

func (e Employee) printInfo() {
	fmt.Printf("Employee ID: %d\n", e.EmployeeID)
	fmt.Printf("Employee Age: %d\n", e.Person.Age)
	fmt.Printf("Employee Name: %d\n", e.EmployeeID)
}
