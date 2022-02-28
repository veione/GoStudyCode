package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	t := head
	addIndex := 0
	for l1 != nil && l2 != nil {
		value := l1.Val + l2.Val
		t.Next = &ListNode{Val: (value + addIndex) % 10}
		addIndex = (value + addIndex) / 10
		t = t.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		t.Next = &ListNode{Val: (l1.Val + addIndex) % 10}
		addIndex = (l1.Val + addIndex) / 10
		t = t.Next
		l1 = l1.Next
	}
	for l2 != nil {
		t.Next = &ListNode{Val: (l2.Val + addIndex) % 10}
		addIndex = (l2.Val + addIndex) / 10
		t = t.Next
		l2 = l2.Next
	}
	if addIndex != 0 {
		t.Next = &ListNode{Val: addIndex}
	}
	return head.Next
}
