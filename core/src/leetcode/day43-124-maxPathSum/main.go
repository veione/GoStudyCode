package day43_124_maxPathSum

import "math"

package day43_543_diameterOfBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		if left < 0 {
			left = 0
		}
		if right < 0 {
			right = 0
		}
		ans = max(ans, left+right+root.Val)
		return max(left,right)	+root.Val
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

