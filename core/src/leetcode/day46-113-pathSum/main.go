package day46_113_pathSum

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := make([][]int, 0, 10)
	var dfs func(root *TreeNode, sum int, path *[]int)
	dfs = func(root *TreeNode, sum int, path *[]int) {
		if root == nil {
			return
		}
		*path = append(*path, root.Val)
		sum += root.Val
		if root.Right == nil && root.Left == nil {
			if sum == targetSum {
				res = append(res, append([]int{}, *path...))
			}
		}
		dfs(root.Left, sum, path)
		dfs(root.Right, sum, path)
		*path = (*path)[:len(*path)-1]
	}
	path := make([]int, 0, 10)
	dfs(root, 0, &path)
	return res
}
