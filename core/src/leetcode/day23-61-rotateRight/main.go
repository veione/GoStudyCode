package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	var oldTail *ListNode
	n := 0
	virNode := &ListNode{}
	virNode.Next = head
	for t := virNode; t != nil; t = t.Next {
		if t.Next == nil {
			oldTail = t
		} else {
			n++
		}
	}
	oldTail.Next = head
	newTail := head
	for i := 1; i <= n-(k%n)-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil
	return newHead
}

func main() {

}
