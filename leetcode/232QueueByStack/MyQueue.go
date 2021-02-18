package main

import "container/list"

type MyQueue struct {
	l1 *list.List
	l2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{list.New(), list.New()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	if this.l1.Len() > 0{
		this.l2 = list.New()
		for ; this.l1.Len() > 0;  {
			this.l2.PushBack(this.l1.Back())
			this.l1.Remove(this.l1.Back())
		}
	}
	this.l1.PushBack(x)

	for ; this.l2.Len() >0 ; {
		this.l1.PushBack(this.l2.Back())
		this.l2.Remove(this.l2.Back())
	}

}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	return this.l1.Remove(this.l1.Front()).(int)
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.l1.Front().Value.(int)
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.l1.Len() == 0
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
