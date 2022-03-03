package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	root := buildTree(head, nil)
	return root
}

func findMid(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := findMid(left, right)
	if mid == nil {
		return nil
	}
	root := &TreeNode{Val: mid.Val}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}

func main() {
	head := &ListNode{Val: -1}
	cur := head
	for i := 1; i <= 6; i++ {
		t := &ListNode{Val: i}
		cur.Next = t
		cur = cur.Next
	}
	buildTree(head.Next, nil)

}
