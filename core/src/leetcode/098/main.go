package main

import "container/list"

Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	lt  list.List
}

func (st *stack) push(val int) {
	element := &list.Element{
		Value: val,
	}
	st.lt.InsertAfter(st.lt.Front(), element)
}


func (st *stack) pop() int {


}
func isValidBST(root *TreeNode) bool {

}

func main() {

}