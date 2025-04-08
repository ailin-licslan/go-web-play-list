package main

func doubleValue(arr *[]int) []int {
	for i := 0; i < len(*arr); i++ {
		(*arr)[i] *= 2
	}
	return *arr
}

func test(p *int) int {
	return *p + 10
}

func swap(a int, b int) {
	var temp int
	temp = a
	a = b
	b = temp
}

// 指针传递
func swap2(a *int, b *int) {
	var temp int
	temp = *a
	*a = *b
	*b = temp
}
