package main

import "fmt"

// 发送1-10 到channel 中
func generator(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

func printer(ch chan int) {
	for num := range ch {
		fmt.Println(num)
	}
}

// 只写通道  chan<- 写道channel
func producer(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
}

// 只读   <-chan 从channel中读
func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println(num)
	}

}
