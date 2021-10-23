package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	ans := 0
	if root == nil {
		return ans
	}
	dfs(root, &ans, root.Val)
	return ans
}

func dfs(root *TreeNode, ans *int, curSum int) {
	if root.Left == nil && root.Right == nil {
		*ans += curSum
		return
	}
	if root.Left != nil {
		dfs(root.Left, ans, curSum*10+root.Left.Val)
	}
	if root.Right != nil {
		dfs(root.Right, ans, curSum*10+root.Right.Val)
	}
}

func main() {
	l := &TreeNode{Val: 2}
	r := &TreeNode{Val: 3}
	t := &TreeNode{Val: 1, Left: l, Right: r}
	ans := sumNumbers(t)
	fmt.Println(ans)
}
