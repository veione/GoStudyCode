package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	small := &ListNode{}
	sm := small
	large := &ListNode{}
	la := large
	for head != nil {
		if head.Val < x {
			sm.Next = head
			sm = sm.Next
		} else {
			la.Next = head
			la = la.Next
		}
		head = head.Next
	}
	la.Next = nil
	sm.Next = large.Next
	return small.Next
}
