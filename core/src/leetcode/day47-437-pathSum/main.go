package day47_437_pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(root *TreeNode, targetSum int) (res int)
	dfs = func(root *TreeNode, targetSum int) (res int) {
		if root == nil {
			return
		}
		res = nodeSum(root, targetSum)
		res += dfs(root.Left, targetSum)
		res += dfs(root.Right, targetSum)
		return
	}
	res := dfs(root, targetSum)
	return res
}

func nodeSum(root *TreeNode, target int) (res int) {
	if root == nil {
		return
	}
	target -= root.Val
	if target == 0 {
		res++
	}
	res += nodeSum(root.Left, target)
	res += nodeSum(root.Right, target)
	return
}
