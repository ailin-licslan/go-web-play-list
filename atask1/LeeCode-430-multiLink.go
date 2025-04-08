package main

/**
430. 扁平化多级双向链表：多级双向链表中，除了指向下一个节点和前一个节点指针之外，它还有一个子链表指针，可能指向单独的双向链表。
这些子列表也可能会有一个或多个自己的子项，依此类推，生成多级数据结构，如下面的示例所示。给定位于列表第一级的头节点，请扁平化列表，
即将这样的多级双向链表展平成普通的双向链表，使所有结点出现在单级双链表中。可以定义一个结构体来表示链表节点，包含 val、prev、next 和 child 指针，
然后使用递归的方法来扁平化链表，先处理当前节点的子链表，再将子链表插入到当前节点和下一个节点之间。
*/

// Node 定义链表节点结构
type Node struct {
	Val   int
	Prev  *Node //prev node
	Next  *Node //next node
	Child *Node //child node
}

func flattenMultiLink(head *Node) *Node {
	if head == nil {
		return head
	}

	//使用stack辅助遍历
	var stack []*Node
	//初始化一个虚拟头节点
	dummy := &Node{}
	//prev指针初始指向 虚拟头节点
	prev := dummy

	//head 入栈
	stack = append(stack, head)

	//遍历时判断 stack 不为空 len(stack) > 0
	for len(stack) > 0 {
		//取出栈顶元素
		curr := stack[len(stack)-1]
		//将栈顶元素从栈中(切片)移除
		stack = stack[:len(stack)-1]
		//将perv的next指向curr 栈顶元素
		prev.Next = curr //(dummy的next是当前的curr栈顶元素)
		//将curr的prev指向prev
		curr.Prev = prev //当前栈顶前一个指针指向dummy(prev)

		//如果curr有next节点 则入栈
		if curr.Next != nil {
			stack = append(stack, curr.Next)
		}
		//如果curr有child节点 则也入栈 并将curr的child置空
		if curr.Child != nil {
			stack = append(stack, curr.Child)
			curr.Child = nil
		}

		//更新prev为curr
		prev = curr
	}

	dummy.Next.Prev = nil
	return dummy.Next
}

// 构建一个多级双向链表示例
func buildSampleList() *Node {

	// 创建节点
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node4 := &Node{Val: 4}
	node5 := &Node{Val: 5}
	node6 := &Node{Val: 6}
	node7 := &Node{Val: 7}
	node8 := &Node{Val: 8}
	node9 := &Node{Val: 9}
	node10 := &Node{Val: 10}

	// 构建链表关系
	node1.Next = node2
	node2.Prev = node1
	node2.Next = node3
	node3.Prev = node2
	node3.Next = node4
	node4.Prev = node3
	node4.Next = node5
	node5.Prev = node4
	node5.Next = node6
	node6.Prev = node5

	node3.Child = node7
	node7.Prev = node3
	node7.Next = node8
	node8.Prev = node7
	node8.Next = node9
	node9.Prev = node8
	node9.Next = node10
	node10.Prev = node9

	return node1
}
