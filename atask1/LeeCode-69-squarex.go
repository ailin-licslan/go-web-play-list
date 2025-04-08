package main

/**
69. x 的平方根：实现 int sqrt(int x) 函数。计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，
小数部分将被舍去。可以使用二分查找法来解决，定义左右边界 left 和 right，然后通过 while 循环不断更新中间值 mid，直到找到满足条件
的平方根或者确定不存在精确的平方根。
*/

// squareOfX
func squareOfX(x int) int {

	//定义左右边界
	left, right := 0, x
	for left < right {
		//计算中间值mid 并防止整数溢出情况
		mid := left + (right-left)/2
		square := mid * mid
		if square == x { //刚刚好找到了
			return mid
		} else if square > x { //说明 x / right那边太大了
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return right
}
