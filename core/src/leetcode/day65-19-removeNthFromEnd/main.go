package day65_19_removeNthFromEnd

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	virturalHead := &ListNode{Next: head}
	left, right := virturalHead, virturalHead
	for n > 0 && right != nil {
		right = right.Next
		n--
	}
	for right.Next != nil {
		right = right.Next
		left = left.Next
	}
	left.Next = left.Next.Next
	return virturalHead.Next
}
