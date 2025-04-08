package main

/**
26. 删除有序数组中的重复项：给你一个有序数组 nums ，请你原地删除重复出现的元素，使每个元素只出现一次，返回删除后数组的新长度。
不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。可以使用双指针法，一个慢指针 i 用于记录不重复元素的位置，
一个快指针 j 用于遍历数组，当 nums[i] 与 nums[j] 不相等时，将 nums[j] 赋值给 nums[i + 1]，并将 i 后移一位。
*/

// DelDup delDup  删除有序数组中的重复项
func DelDup(arr []int) int {

	if len(arr) == 0 {
		return 0
	}
	slow := 0 //慢指针
	fast := 0 //快指针

	for fast < len(arr) {
		//如果快指针指向的值 和 慢指针指向的值不一致
		if arr[fast] != arr[slow] {
			//慢指针向前移动一位
			slow++
			//维护arr[0...slow]无重复  并把fast指针此时指向的值赋值给slow指针所在的位置
			arr[slow] = arr[fast]
		}
		//快指针继续向前移动
		fast++
	}

	for i := 0; i < len(arr); i++ {
		println("value is :", arr[i])
	}
	return slow + 1

}
