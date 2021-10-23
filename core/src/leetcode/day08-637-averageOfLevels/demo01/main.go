package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := make([]float64, 0, 10)
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		sum := 0
		for i := 0; i < size; i++ {
			top := queue[0]
			queue = queue[1:]
			sum += top.Val
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right !=nil {
				queue = append(queue, top.Right)
			}
		}
		res = append(res, float64(sum)/float64(size))
	}
	return res
}
