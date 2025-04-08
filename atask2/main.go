package main

import (
	"fmt"
	"sync"
)

/*
*
a xxxx  每个名字加 a 为了main放最下面
*/
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

	//TEST Exercise 4 CHANNEL
	ch := make(chan int)
	go generator(ch) //ch 里面填充数据
	go printer(ch)   //打印ch通道中的数据
	var input string
	_, _ = fmt.Scanln(&input)

	//带缓冲区
	chBuffer := make(chan int, 20)
	go producer(chBuffer)
	go consumer(chBuffer)
	var inputBuffer string
	_, _ = fmt.Scanln(&inputBuffer)

	//TEST Exercise 5 锁机制
	var wgs sync.WaitGroup
	wgs.Add(10) //设置需要等待的协程数为10个
	for i := 0; i < 10; i++ {
		//启动10个协程 执行increment方法
		go increment(&wgs)
	}
	wgs.Wait() //阻塞main 函数 直到10个协程都完成计数器递增操作
	fmt.Println("Final counter value:", counter)

	var wgs2 sync.WaitGroup
	wgs2.Add(10) //设置需要等待的协程数为10个
	for i := 0; i < 10; i++ {
		//启动10个协程 执行increment方法
		go increment2(&wgs2)
	}
	wgs2.Wait() //阻塞main 函数 直到10个协程都完成计数器递增操作
	fmt.Println("Final counter2 value:", counter2)
}
