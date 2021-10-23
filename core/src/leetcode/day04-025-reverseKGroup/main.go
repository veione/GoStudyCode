package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hear := &ListNode{Next: head}
	pre := hear
	for head != nil {
		tail := pre
		for i:=0; i< k ;i++ {
			tail = tail.Next
			if tail == nil {
				return hear.Next
			}
		}
		temp := tail.Next
		header, tailer := reverse(head,tail)
		tailer.Next = temp
		pre.Next = header
		head = tailer.Next
		pre = tailer
	}
	return hear.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		nex := p.Next
		p.Next = prev
		prev = p
		p = nex
	}
	return tail, head
}

func main() {

}