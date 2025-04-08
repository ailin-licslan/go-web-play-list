package main

import (
	"sort"
)

/**
56. 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。可以先对区间数组按照区间的起始位置进行排序，
然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，
则将当前区间添加到切片中。
*/

// mergeRange  这个不太理解了 断点一点点调试可以理解了
func mergeRange(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	//先排序一下  排序依据是区间的起始位置 intervals[i][0] 这样排序后 重叠的区间会相邻
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//第一个区间的内容添加到结果切片result中
	var result [][]int
	result = append(result, intervals[0]) //最小的加进去 [1,3]  这个时候按照每个元素[start,end] 中start 依次升序排序了  1,3  2,6  3,5  8,10  15,18

	//从第二个区间index =1开始遍历 intervals 数组
	for i := 1; i < len(intervals); i++ {

		current := intervals[i]       //[2,6]  start from i=1
		last := result[len(result)-1] //[1,3]  start from i=1

		//判断current 和 last 是否重叠 检查current 的起始位置 current[0] 是否<= last的结束位置 last[1]
		if current[0] <= last[1] { // 2 < 3
			//重叠的话 更新last的结束位置  选择相对大的(last / current)
			last[1] = maxOne(last[1], current[1]) //这个时候就把 6挑出来作为 last[1]了 因为  6 > 3 选择较大的那个
		} else {
			//不重叠 直接将 current添加到result中
			result = append(result, current) //不重叠的话较大的依次往后面加
		}
	}

	return result
}

func maxOne(a, b int) int {
	if a > b {
		return a
	}
	return b
}
