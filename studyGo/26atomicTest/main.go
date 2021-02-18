package main

import (
	"fmt"
	"sync"
	"sync/atomic"

)

var total int32

func worker(wg *sync.WaitGroup ){
	defer wg.Done()
	var i int32
	for i = 0; i < 10; i++ {
		atomic.AddInt32(&total, i)
		fmt.Println("tatal:",total)
	}
}

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)

	wg.Wait()
	fmt.Println("main: ",total)
}
