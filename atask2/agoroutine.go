package main

import "time"

func printOdd() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			println("Odd: ", i)
		}
	}
}

// printEven
func printEven() {
	for i := 2; i <= 10; i++ {
		if i%2 == 0 {
			time.Sleep(1 * time.Second)
			println("Even: ", i)
		}
	}
}
