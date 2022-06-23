package day43_687_longestUnivaluePath

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ansLeft := 0
		ansRight := 0
		if root.Left != nil && root.Left.Val == root.Val {
			ansLeft = left + 1
		}
		if root.Right != nil && root.Right.Val == root.Val {
			ansRight = right + 1
		}
		ans = max(ans, ansLeft+ansRight)
		return max(ansLeft, ansRight)
	}
	dfs(root)
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
