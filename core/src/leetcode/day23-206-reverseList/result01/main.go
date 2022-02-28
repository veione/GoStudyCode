package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var virtualNode *ListNode
	l, r := virtualNode, head
	for r != nil {
		t := r.Next
		r.Next = l
		l = r
		r = t
	}
	return l
}
