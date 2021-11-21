package main

import "fmt"

type LRUCache struct {
	head, tail *element
	m          map[int]*element
	cap        int
	length     int
}

type element struct {
	k, v      int
	pre, next *element
}

func newElement(k, v int)*element {
	return &element{
		k:    k,
		v:    v,
		pre:  nil,
		next: nil,
	}
}

func Constructor(capacity int) LRUCache {
	res := LRUCache{
		head: newElement(0,0),
		tail: newElement(0,0),
		m:   make(map[int]*element, capacity),
		cap: capacity,
		length: 0,
	}
	res.head.next = res.tail
	res.tail.pre = res.head
	return res
}

func (this *LRUCache) Get(key int) int {
	if element, ok := this.m[key]; ok {
		this.MoveToFront(element)
		return element.v
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if element, ok := this.m[key]; ok {
		element.v = value
		this.MoveToFront(element)
		return
	}
	if this.length >= this.cap {
		this.DelTail()
	}
	e := &element{
		k:    key,
		v:    value,
		pre:  nil,
		next: nil,
	}
	this.AddToFront(e)
}

func (this *LRUCache) DelElement (e *element) {
	preNode := e.pre
	nextNode := e.next
	preNode.next = nextNode
	nextNode.pre = preNode
	delete(this.m, e.k)
	this.length --
}


func (this *LRUCache) MoveToFront(e *element){
	this.DelElement(e)
	this.AddToFront(e)
}

func (this *LRUCache) DelTail() {
	this.DelElement(this.tail.pre)
}

func (this *LRUCache) AddToFront(e *element) {
	oldHead := this.head.next
	e.next=oldHead
	oldHead.pre = e
	this.head.next = e
	e.pre = this.head
	this.m[e.k] = e
	this.length ++
}

func (this *LRUCache) PrintList(){
	for e := this.head.next; e != this.tail; e = e.next {
		fmt.Printf("key:%v  value:%v ", e.k, e.v)
	}
	fmt.Println()
}



func main()  {
	lRUCache := Constructor(2);
	lRUCache.Put(2, 1); // 缓存是 {1=1}
	lRUCache.PrintList()
	lRUCache.Put(1, 1); // 缓存是 {1=1, 2=2}
	lRUCache.PrintList()
	lRUCache.Put(2, 3); // 缓存是 {2=3, 2=1}
	lRUCache.PrintList()
	lRUCache.Put(4, 1); // 缓存是 {4=1, 2=3}
	lRUCache.PrintList()
	lRUCache.Get(1);    // 返回 -1
	lRUCache.PrintList()
	//lRUCache.Put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	//lRUCache.printList()
	//lRUCache.Get(2);    // 返回 -1 (未找到)
	//lRUCache.printList()
	//lRUCache.Put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	//lRUCache.printList()
	lRUCache.Get(2);    // 返回 -1 (未找到)
	lRUCache.PrintList()
	//lRUCache.Get(3);    // 返回 3
	//lRUCache.printList()
	//lRUCache.Get(4);    // 返回 4
	//lRUCache.printList()

}
