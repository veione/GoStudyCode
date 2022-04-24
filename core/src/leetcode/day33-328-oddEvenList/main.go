package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	cur1 := head
	head2 := head.Next
	cur2 := head2
	for cur1 != nil && cur2 != nil {
		if cur1.Next != nil {
			cur1.Next = cur1.Next.Next
			if cur1.Next != nil {
				cur1 = cur1.Next
			}
		}
		if cur2.Next != nil {
			cur2.Next = cur2.Next.Next
			cur2 = cur2.Next
		}
	}
	cur1.Next = head2
	return head
}

func main() {
	head := &ListNode{}
	datas := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var preNode *ListNode = head
	for i := range datas {
		node := &ListNode{Val: datas[i]}
		preNode.Next = node
		preNode = preNode.Next
	}

}
