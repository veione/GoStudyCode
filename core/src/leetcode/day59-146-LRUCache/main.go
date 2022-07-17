package main

import (
	"container/list"
	"fmt"
)

type kv struct {
	key int
	val int
}

type LRUCache struct {
	mapData  map[int]*list.Element
	listData *list.List
	cap      int
}

func Constructor(capacity int) LRUCache {
	lruCache := LRUCache{
		cap:      capacity,
		listData: list.New(),
		mapData:  make(map[int]*list.Element),
	}
	return lruCache
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.mapData[key]; !ok {
		return -1
	} else {
		this.listData.MoveToFront(val)
		return val.Value.(*kv).val
	}
}

func (this *LRUCache) Put(key int, value int) {
	if this.cap <= 0 {
		return
	}
	if element, ok := this.mapData[key]; ok {
		element.Value.(*kv).val = value
		this.listData.MoveToFront(element)
		return
	}
	e := this.listData.PushFront(&kv{key: key, val: value})
	this.mapData[key] = e
	if len(this.mapData) > this.cap {
		back := this.listData.Back()
		delete(this.mapData, back.Value.(*kv).key)
		this.listData.Remove(back)
	}
}

func (this *LRUCache) print() {
	for index := this.listData.Front(); index != nil; index = index.Next() {
		fmt.Printf(" %v", index.Value.(int))
	}
}

func main() {
	cache := Constructor(3)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	cache.Put(3, 3)
	cache.Put(4, 4)
	cache.Put(3, -3)
	cache.print()
}
