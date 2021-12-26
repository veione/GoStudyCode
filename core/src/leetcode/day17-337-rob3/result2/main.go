package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	f, g := dfs(root)
	return max(f, g)
}

func dfs(node *TreeNode)(f, g int) {
	if node == nil {
		return 0 ,0
	}
	lf,lg := dfs(node.Left)
	rf,rg := dfs(node.Right)
	f = lg + rg + node.Val
	g = max(lf, lg) + max(rf, rg)
	return f, g
}


func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}
