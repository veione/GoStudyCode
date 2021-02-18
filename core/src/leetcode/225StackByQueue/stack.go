package main

import "container/list"

type MyStack struct {
	l1 *list.List
	l2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		list.New(),
		list.New(),
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.l2.PushBack(x)
	if this.l1.Len() > 0 {
		for e := this.l1.Front(); e!= nil; e = e.Next() {
			this.l2.PushBack(e.Value)
		}
	}
	this.l1 = this.l2
	this.l2 = list.New()
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.l1.Remove(this.l1.Front()).(int)

}


/** Get the top element. */
func (this *MyStack) Top() int {
	return this.l1.Front().Value.(int)
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l1.Len() == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
