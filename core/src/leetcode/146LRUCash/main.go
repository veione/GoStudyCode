package main

import (
	"container/list"
	"fmt"
)

type item struct {
	key int
	val int
}

type LRUCache struct {
	list *list.List
	capacity int
}


func Constructor(capacity int) LRUCache {
	return LRUCache{list.New(), capacity}
}


func (this *LRUCache) Get(key int) int {
	var res int = -1
	e := this.list.Front()
	for ; e != nil ; e = e.Next() {
		if t, ok := e.Value.(*item); ok {
			if t.key == key {
				res = t.val
				break
			}
		}
	}
	if res != -1{
		this.list.MoveToFront(e)
	}
	return res
}


func (this *LRUCache) Put(key int, value int)  {
	change := false
	e := this.list.Front()
	for ; e != nil ; e = e.Next() {
		if t, ok := e.Value.(*item); ok {
			if t.key == key {
				t.val = value
				change = true
				break
			}
		}
	}
	if !change {
		if this.list.Len() == this.capacity{
			this.list.Remove(this.list.Back())
		}
		this.list.PushFront(&item{key,value})
	} else {
		this.list.MoveToFront(e)
	}
}

func (this *LRUCache) printList(){
	for e:= this.list.Front(); e != nil; e = e.Next(){
		if ite, ok := e.Value.(*item); ok {
			fmt.Printf("key:%d  value:%d ->",ite.key, ite.val)
		}
	}
	fmt.Println()
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main()  {
	lRUCache := Constructor(2);
	lRUCache.Put(2, 1); // 缓存是 {1=1}
	lRUCache.printList()
	lRUCache.Put(1, 1); // 缓存是 {1=1, 2=2}
	lRUCache.printList()
	lRUCache.Put(2, 3); // 缓存是 {1=1, 2=2}
	lRUCache.printList()
	lRUCache.Put(4, 1); // 缓存是 {1=1, 2=2}
	lRUCache.printList()
	lRUCache.Get(1);    // 返回 1
	lRUCache.printList()
	//lRUCache.Put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//lRUCache.printList()
	//lRUCache.Get(2);    // 返回 -1 (未找到)
	//lRUCache.printList()
	//lRUCache.Put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//lRUCache.printList()
	lRUCache.Get(2);    // 返回 -1 (未找到)
	lRUCache.printList()
	//lRUCache.Get(3);    // 返回 3
	//lRUCache.printList()
	//lRUCache.Get(4);    // 返回 4
	//lRUCache.printList()

}