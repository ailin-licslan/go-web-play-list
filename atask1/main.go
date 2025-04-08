package main

import "fmt"

func main() {

	//TEST 136
	arr1 := []int{1, 5, 2, 1, 2}
	FindNumOccurOnce(arr1) //exp ===> 5

	//TEST 26
	arr2 := []int{1, 2, 2, 3, 3, 5, 6}      // 1 2 3 5 6
	println("The result is ", DelDup(arr2)) //exp ===> 5

	//TEST 21
	a1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	a2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	merge2Link(a1, a2)

	//TEST 344
	s := []byte("abc")
	println("The result is :", ReverseStr(s)) // exp ===> cba

	//TEST 69
	a := 25
	println("The squareOfX result is :", squareOfX(a)) //exp ===> 5

	//TEST 46
	nums := []int{1, 2, 3} //[1 2 3][1 3 2]  [2 1 3][2 3 1]  [3 1 2][3 2 1]
	result := fullPermutation(nums)
	for _, num := range result {
		fmt.Println(num)
	}

	//TEST 56
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}, {3, 5}}
	merged := mergeRange(intervals)
	fmt.Println(merged) // exp ===> [[1 6] [8 10] [15 18]]

	//TEST 430
	head := buildSampleList()
	newHead := flattenMultiLink(head)
	// 遍历扁平化后的链表
	for p := newHead; p != nil; p = p.Next {
		fmt.Printf("%d ", p.Val) // exp ===> 1 2 3 7 8 9 10 4 5 6
	}

	//TEST 198
	numbs := []int{1, 2, 3, 1}
	data := houseRobber(numbs)
	fmt.Println("Maximum amount that can be robbed: ", data)

	//TEST 729
	cal := Constructor()
	fmt.Println(cal.book(10, 20)) // .....is ok to book true
	fmt.Println(cal.book(15, 25)) // .....can't book false
	fmt.Println(cal.book(20, 30)) // .....is ok to book true

}
