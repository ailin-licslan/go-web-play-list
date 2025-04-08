package main

import (
	"sync"
	"sync/atomic"
)

// 定义共享计数器和互斥锁
var (
	counter2 int64
	counter  int        //定义共享资源  counter是一个共享的整数类型的计数器，初始值为 0。
	mu       sync.Mutex //互斥锁：mu 是一个 sync.Mutex 类型的互斥锁，用于保护对 counter 的访问
)

// 接受一个指向 sync.WaitGroup 的指针 wg 作为参数，sync.WaitGroup 用于等待所有协程完成。
func increment(wg *sync.WaitGroup) {
	// 使用 defer wg.Done() 确保在函数结束时通知 sync.WaitGroup 该协程已完成
	defer wg.Done()
	// 在 for 循环中，每次对 counter 进行递增操作前，调用 mu.Lock() 锁定互斥锁，防止其他协程同时访问 counter；
	//递增操作完成后，调用 mu.Unlock() 解锁互斥锁，允许其他协程访问。
	for i := 0; i < 1000; i++ {
		//保证多协程下对共享资源的安全访问  避免数据竞争问题
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func increment2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		//atomic.AddInt64 函数会保证对counter2的递增操作是原子的 不会被其他的协程操作所打断 实现了无锁的计数器 和 Java 里面很像
		atomic.AddInt64(&counter2, 1)
	}

}
