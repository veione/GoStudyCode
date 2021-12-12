package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    virtualHead := &ListNode{}
    curNode := virtualHead
    h1, h2 := l1, l2
    carry := 0
    for h1 != nil || h2 != nil || carry != 0 {
    	sum := 0
    	if h1 != nil {
    		sum += h1.Val
    		h1 = h1.Next
		}
		if h2 != nil {
			sum += h2.Val
			h2 = h2.Next
		}
		if carry != 0 {
			sum += carry
		}
		curNode.Next = &ListNode{
			Val:  sum%10,
		}
		carry = sum / 10
		curNode = curNode.Next
	}
	return virtualHead.Next
}

func reverseList(l *ListNode) *ListNode {
	cur := l
	var pre *ListNode = nil
	next := l.Next
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur =next
	}
	return pre
}