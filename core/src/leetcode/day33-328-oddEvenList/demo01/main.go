package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curOdd := head
	evenHead := head.Next
	curEven := evenHead
	for curEven != nil && curEven.Next != nil {
		curOdd.Next = curEven.Next
		curOdd = curOdd.Next
		curEven.Next = curOdd.Next
		curEven = curEven.Next
	}
	curOdd.Next = evenHead
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
	for node := head.Next; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println("转换后")
	newHead := oddEvenList(head.Next)
	for node := newHead; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
}
