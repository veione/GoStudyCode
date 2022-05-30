package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
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
	return head, evenHead
}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre
}

func mergeList(head1, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}
	if head2 == nil {
		return head1
	}
	var head *ListNode
	cur1 := head1
	cur2 := head2
	if head1.Val < head2.Val {
		head = head1
		cur1 = cur1.Next
	} else {
		head = head2
		cur2 = cur2.Next
	}
	cur := head
	for cur1 != nil && cur2 != nil {
		if cur1.Val < cur2.Val {
			cur.Next = cur1
			cur1 = cur1.Next
		} else {
			cur.Next = cur2
			cur2 = cur2.Next
		}
		cur = cur.Next
	}
	if cur1 != nil {
		cur.Next = cur1
	}
	if cur2 != nil {
		cur.Next = cur2
	}
	return head
}

func main() {
	head := &ListNode{}
	datas := []int{1, 8, 3, 6, 5, 4, 7, 2, 9}
	var preNode *ListNode = head
	for i := range datas {
		node := &ListNode{Val: datas[i]}
		preNode.Next = node
		preNode = preNode.Next
	}
	fmt.Print("原始链表:")
	for node := head.Next; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println()
	fmt.Println("奇数偶数拆分后:")
	oddHead, evenHead := oddEvenList(head.Next)

	fmt.Print("奇数链表:")
	for node := oddHead; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println()
	fmt.Print("偶数链表")
	for node := evenHead; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println()

	fmt.Print("翻转偶链表:")
	reverseHead := reverseList(evenHead)
	for node := reverseHead; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
	fmt.Println()

	res := mergeList(oddHead, reverseHead)
	fmt.Print("合并奇偶链表后")
	for node := res; node != nil; node = node.Next {
		fmt.Printf("%v ", node.Val)
	}
}
