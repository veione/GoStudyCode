package main

import (
	"fmt"
	"sync"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var sum int
	for i:= 0; i < 10; i++ {
		go func(k int) {
			cond.L.Lock()
			sum ++
			fmt.Printf("%v号线程执行，sum:%v \n", k, sum)
			cond.L.Unlock()
			cond.Broadcast()
		}(i)
	}
	fmt.Println("德玛西亚")
	cond.L.Lock()
	for sum != 10{
		fmt.Println("主线程被唤醒 一次")
		cond.Wait()
	}
	cond.L.Unlock()
	fmt.Println("主线程执行完成")
}
