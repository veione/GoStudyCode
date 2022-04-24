package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur := head
	res := &ListNode{}
	for cur != nil {
		next := cur.Next
		inserSortedtList(res, cur)
		cur = next
	}
	return res.Next
}

func inserSortedtList(head *ListNode, node *ListNode) {
	for head.Next != nil && node.Val > head.Next.Val {
		head = head.Next
	}
	next := head.Next
	head.Next = node
	node.Next = next
}
