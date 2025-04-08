package main

/**
21. 合并两个有序链表：将两个升序链表合并为一个新的升序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。可以定义一个函数，
接收两个链表的头节点作为参数，在函数内部使用双指针法，通过比较两个链表节点的值，将较小值的节点添加到新链表中，直到其中一个链表为空，
然后将另一个链表剩余的节点添加到新链表中。
*/

// ListNode 定义node
type ListNode struct {
	Val  int
	Next *ListNode
}

// merge2Link  合并2个升序链表  双指针
func merge2Link(a1 *ListNode, a2 *ListNode) *ListNode {

	dummy := &ListNode{}
	cur := dummy

	//遍历过程中2个链表都不为空
	for a1 != nil && a2 != nil {
		if a1.Val < a2.Val { //如果a1链表的元素遍历时 比 a2 小
			cur.Next = a1 //把 a1的 挂在 新链表 cur上
			a1 = a1.Next
		} else {
			cur.Next = a2 //反之把 a2的 挂在 新链表 cur上
			a2 = a2.Next
		}
		cur = cur.Next //新链表也往前移动一下
	}

	//有个先遍历完 为空时  2个链表长度不同时  一般长的没有遍历完有剩余的  剩余 的节点直接接在新链表上面去
	if a1 != nil {
		cur.Next = a1
	} else {
		cur.Next = a2
	}

	//遍历打印新链表
	for dummy.Next != nil {
		println("The order is: ", dummy.Val)
		dummy = dummy.Next
	}

	return dummy.Next

}
