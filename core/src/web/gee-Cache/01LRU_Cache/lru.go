package lru

import "container/list"

// Cache is a LRU cache. It is not safe for concurrent access.
type Cache struct {
	maxBytes int64
	nbytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	// optional and executed when an entry is purged.
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

// Value use Len to count how many bytes it takes
type Value interface {
	Len() int
}

func NewCache(maxBytes int64, onEvicted func(string, Value))* Cache  {
	return &Cache{
		maxBytes:  maxBytes,
		nbytes:    0,
		ll:        list.New(),
		cache:     make(map[string]*list.Element),
		OnEvicted: onEvicted,
	}
}

func (this *Cache)Get(key string)(value Value, ok bool){
	if value,ok := this.cache[key]; ok{
		entry := value.Value.(*entry)
		this.ll.MoveToFront(value)
		return entry.value, true
	}
	return
}

func (this *Cache)RemoveOldest()  {
	e := this.ll.Back()
	if e != nil {
		this.ll.Remove(e)
		kv := e.Value.(*entry)
		delete(this.cache, kv.key)
		this.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if this.OnEvicted !=nil{
			this.OnEvicted(kv.key, kv.value)
		}
	}
}

func (this *Cache)Add(key string, value Value) {
	e := entry{
		key:   key,
		value: value,
	}
	if node, ok := this.cache[key]; !ok {
		ele := &list.Element{
			Value: e,
		}
		this.ll.InsertBefore(ele, this.ll.Front())
		this.cache[key]= ele
		this.nbytes += int64(len(key)) + int64(value.Len())
	}else {
		kv := node.Value.(entry)
		kv.value = value
		this.ll.MoveToFront(node)
		this.nbytes += int64(len(key)) + int64(kv.value.Len())
	}

	if this.maxBytes >0 && this.nbytes >= this.maxBytes {
		this.RemoveOldest()
	}
}
