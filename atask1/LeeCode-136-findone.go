package main

/**
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/

// FindNumOccurOnce setMap
func FindNumOccurOnce(arr []int) int {

	myMap := make(map[int]int)

	//遍历数组 1种写法
	for i := 0; i < len(arr); i++ {
		println("The element is: ", arr[i])
	}

	//遍历数组 2种写法
	for _, value := range arr {
		//如果相同 map的value +1
		myMap[value] = myMap[value] + 1
		//myMap[value]++
	}

	//找出出现次数为1的元素  value ==1
	for key, value := range myMap {
		if value == 1 {
			println("The only one element is: ", key)
			return key
		}
	}

	//没有找到的返回-1
	return -1

}
