package main

import (
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	virtualHead := &ListNode{Next: head, Val: math.MinInt32}
	lastSortNode := head
	cur:= head.Next
	for cur != nil {
		if lastSortNode.Val <= cur.Val {
			lastSortNode = lastSortNode.Next
		}else {
			pre := virtualHead
			for pre.Next != nil && pre.Next.Val <= cur.Val {
				pre = pre.Next
			}
			lastSortNode.Next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
		}
		cur = lastSortNode.Next
	}
	return virtualHead.Next
}


