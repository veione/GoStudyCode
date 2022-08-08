package result1

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 10)
	stack = append(stack, root)
	var pre *TreeNode
	var cur *TreeNode
	for len(stack) > 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if pre != nil {
			pre.Left = nil
			pre.Right = cur
		}
		pre = cur
		// todo 输出T
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		cur = cur.Left
	}
}
