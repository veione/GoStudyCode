package main

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	mu sync.Mutex
	initialized uint32
}

func (once *Once) Do(f func())  {
	if atomic.LoadUint32(&initialized) == 1{
		return
	}
	mu.Lock()
	defer mu.Unlock()
	if once.initialized == 0 {
		atomic.StoreUint32(&initialized, 1)
		f()
	}
	return
}