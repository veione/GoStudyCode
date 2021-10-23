package myOnce

import (
	"sync"
	"sync/atomic"
)

type Once struct {
	done uint32
	m    sync.Mutex
}

func (once *Once) Do(f func()) {
	if atomic.LoadUint32(&once.done) == 0 {
		once.doSlow(f)
	}
}

func (once *Once) doSlow(f func()) {
	once.m.Lock()
	defer once.m.Unlock()
	if once.done == 0 {
		atomic.StoreUint32(&once.done, 1)
		f()
	}
}
