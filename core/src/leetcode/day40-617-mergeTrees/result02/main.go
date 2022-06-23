package result02

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	queue := make([]*TreeNode, 0, 10)
	queue1 := make([]*TreeNode, 0, 10)
	queue2 := make([]*TreeNode, 0, 10)
	root := &TreeNode{Val: root1.Val + root2.Val}
	queue = append(queue, root)
	queue1 = append(queue1, root1)
	queue2 = append(queue2, root2)
	for len(queue1) > 0 && len(queue2) > 0 {
		node := queue[0]
		queue = queue[1:]
		node1 := queue1[0]
		queue1 = queue1[1:]
		node2 := queue2[0]
		queue2 = queue2[1:]
		if node1.Left != nil || node2.Left != nil {
			if node1.Left != nil && node2.Left != nil {
				node.Left = &TreeNode{Val: node1.Left.Val + node2.Left.Val}
				queue1 = append(queue1, node1.Left)
				queue2 = append(queue2, node2.Left)
				queue = append(queue, node.Left)
			} else if node1.Left != nil {
				node.Left = node1.Left
			} else {
				node.Left = node2.Left
			}
		}

		if node1.Right != nil || node2.Right != nil {
			if node1.Right != nil && node2.Right != nil {
				node.Right = &TreeNode{Val: node1.Right.Val + node2.Right.Val}
				queue1 = append(queue1, node1.Right)
				queue2 = append(queue2, node2.Right)
				queue = append(queue, node.Right)
			} else if node1.Right != nil {
				node.Right = node1.Right
			} else {
				node.Right = node2.Right
			}
		}
	}
	return root
}
