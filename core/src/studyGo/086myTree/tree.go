package _86myTree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// stack 存储 右节点
func PreOrder(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 10)
	stack = append(stack, root)
	var cur *TreeNode
	for len(stack) > 0 || cur != nil {
		if cur == nil {
			cur = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		if cur != nil {
			fmt.Printf("%v ", cur.Val)
		}
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		cur = cur.Left
	}
}

func PreOrder1(root *TreeNode) {
	stack := make([]*TreeNode, 0, 10)
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			fmt.Printf("%v ", cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			t := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur = t.Right
		}
	}
}

func MidleOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0, 10)
	cur := root
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			temp := stack[len(stack)-1]
			fmt.Printf("%v ", temp.Val)
			stack = stack[:len(stack)-1]
			cur = temp.Right
		}
	}
}

// https://bitbrave.github.io/2020/02/18/%E4%BA%8C%E5%8F%89%E6%A0%91%E7%9A%84%E5%89%8D%E5%BA%8F%E3%80%81%E4%B8%AD%E5%BA%8F%E3%80%81%E5%90%8E%E5%BA%8F%E9%81%8D%E5%8E%86%E7%9A%84%E9%9D%9E%E9%80%92%E5%BD%92%E5%AE%9E%E7%8E%B0/
func LastOrder(root *TreeNode) {
	stack := make([]*TreeNode, 0, 10)
	cur := root
	// 记录上一个 访问过得节点
	var last_visit *TreeNode
	for len(stack) > 0 || cur != nil {
		if cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		} else {
			top := stack[len(stack)-1]
			// 如果右节点为空 或者 右节点已访问
			if top.Right == nil || top.Right == last_visit {
				fmt.Printf("%v ", top.Val)
				stack = stack[:len(stack)-1]
				// 记录当前访问的节点
				last_visit = top
				// 重置 cur节点
				cur = nil
			} else {
				cur = top.Right
			}
		}
	}
}

// 层次遍历
func levelOrder(root *TreeNode) [][]int {
	order := make([][]int, 0, 10)
	if root == nil {
		return order
	}
	queue := make([]*TreeNode, 0, 10)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0, size)
		for i := 0; i < size; i++ {
			node := queue[0]
			level = append(level, node.Val)
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		order = append(order, level)
	}
	return order
}

// 前序 中序 构建 二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	node := &TreeNode{Val: preorder[0]}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//node.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[0:i])
	node.Left = buildTree(preorder[1:i+1], inorder[0:i])
	//node.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	node.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return node
}

func createTree() *TreeNode {
	node1 := &TreeNode{1, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node3 := &TreeNode{3, nil, nil}
	node4 := &TreeNode{4, nil, nil}
	node5 := &TreeNode{5, nil, nil}

	node1.Left, node1.Right = node2, node5
	node2.Left, node2.Right = node3, node4

	return node1
}
