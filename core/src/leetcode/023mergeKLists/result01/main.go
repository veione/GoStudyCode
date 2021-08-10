package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	cur := &ListNode{}
	if l1.Val < l2.Val {
		cur = l1
		l1 = l1.Next
	} else {
		cur = l2
		l2 = l2.Next
	}
	head := cur
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	}
	if l2 != nil {
		cur.Next = l2
	}
	return head
}

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	return merge(lists, 0, n -1)
}

func merge(lists []*ListNode, l, r int)*ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	left := merge(lists, l, mid)
	right := merge(lists, mid+1, r)
	return mergeTwoLists(left, right)
}