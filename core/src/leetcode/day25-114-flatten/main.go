package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	list := make([]*TreeNode, 0, 10)
	preorderTraversal(root, &list)
	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}

func preorderTraversal(root *TreeNode, list *[]*TreeNode) {
	if root != nil {
		*list = append(*list, root)
		preorderTraversal(root.Left, list)
		preorderTraversal(root.Right, list)
	}
}
