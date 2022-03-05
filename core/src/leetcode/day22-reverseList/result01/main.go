package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var virtualNode *ListNode = &ListNode{}
	virtualNode.Next = head
	left, right := virtualNode, head
	t := right.Next
	for t != nil {
		right.Next = left
		left = right
		right = t
		t = t.Next
	}
	
	return right
}

func main() {

}
