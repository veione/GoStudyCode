package main

import (
	"sync"
	"sync/atomic"
)

type singleton struct {}

//原子操作配合互斥锁可以实现非常高效的单件模式。互斥锁的代价比普通整数的原子读写高很多，在性能敏感的地方可以增加一个数字型的标志位，通过原子检测标志位状态降低互斥锁的使用次数来提高性能。

var (
	instance *singleton
	mu 		 sync.Mutex
	initialized uint32
)

func Instance() *singleton {
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}
	mu.Lock()
	defer mu.Unlock()
    if instance == nil {
		instance = &singleton{}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}
