package _86myTree

import (
	"fmt"
	"testing"
)

func TestTreePreOrder(t *testing.T) {
	root := createTree()
	PreOrder1(root)
}

func TestTreeMildOrder(t *testing.T) {
	root := createTree()
	MidleOrder(root)
}

func TestTreeLastOrder(t *testing.T) {
	root := createTree()
	LastOrder(root)
}

func TestBuildTree(t *testing.T) {
	var pre = []int{1, 2, 3, 4, 5, 6}
	var in = []int{3, 2, 4, 1, 5, 6}
	root := buildTree(pre, in)
	PreOrder(root)
}

func TestLevelOrder(t *testing.T) {
	root := createTree()
	ans := levelOrder(root)
	for _, level := range ans {
		for _, val := range level {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
