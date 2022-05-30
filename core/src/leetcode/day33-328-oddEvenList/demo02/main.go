package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	old := head
	even := head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		old.Next = even.Next
		old = old.Next
		even.Next = old.Next
		even = even.Next
	}
	old.Next = evenHead
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
