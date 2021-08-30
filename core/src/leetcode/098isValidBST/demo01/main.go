package main

import (
	"container/list"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type stack struct {
	list *list.List
}

func NewStack() *stack {
	return &stack{
		list: list.New(),
	}
}

func (st *stack) push(val interface{}) {
	st.list.PushBack(val)
}

func (st *stack) pop() interface{} {
	element := st.list.Back()
	if element != nil {
		st.list.Remove(element)
		return element.Value
	}
	return nil
}

func (st *stack) isEmpty() bool {
	return st.list.Len() == 0
}

func (st *stack) peek() interface{} {
	element := st.list.Back()
	if element != nil {
		return element.Value
	}
	return nil
}

func isValidBST(root *TreeNode) bool {
	st := NewStack()
	pre := math.MinInt32
	for root != nil || !st.isEmpty(){
		 for root!= nil{
		 	st.push(root)
		 	root = root.Left
		 }
		 if !st.isEmpty(){
		 	top := st.pop().(TreeNode)
		 	if top.Val <= pre{
		 		return false
			}
		 	pre = top.Val
		 	root = top.Right
		 }
	}
	return true
}

func main() {
	st := NewStack()
	st.push(1)
	st.push(2)
	st.push(3)
	st.push(4)
	for !st.isEmpty() {
		top := st.pop()
		fmt.Printf("%v ", top)

	}
	fmt.Println()
	st.push(1)
	st.push(2)
	st.push(3)
	fmt.Printf("%v ", st.pop())

}
