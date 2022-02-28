package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	virNode := &ListNode{}
	virNode.Next = head

	// 左侧链表 尾节点
	leftTail := virNode
	for i := 1; i <= left-1; i++ {
		leftTail = leftTail.Next
	}
	// 待翻转链表 头结点
	midHead := leftTail.Next
	// 右侧链表头结点

	midTail := virNode
	for i := 1; i <= right; i++ {
		midTail = midTail.Next
	}
	rightHead := midTail.Next
	// 断开链表
	midTail.Next = nil
	leftTail.Next = nil

	newHead := reverse(midHead)
	// 连接成新表
	leftTail.Next = newHead
	midHead.Next = rightHead
	return virNode.Next
}
