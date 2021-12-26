package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	f, g := make(map[*TreeNode]int), make(map[*TreeNode]int)
	dfs(root, f, g)
	return max(f[root], g[root])
}

func dfs(root *TreeNode, f, g map[*TreeNode]int) {
	if root == nil {
		return
	}
	dfs(root.Left, f, g)
	dfs(root.Right, f, g)
	f[root] = max(g[root.Left], g[root.Right]) + root.Val
	g[root] = max(g[root.Left], f[root.Left]) + max(g[root.Right], f[root.Right])
}


func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}
