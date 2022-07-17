package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func loop(wg *sync.WaitGroup, index int) {
	for true {

	}
	wg.Done()
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go loop(&wg, i)
	}
	wg.Wait()
	fmt.Println("所有groutine 已开启")
}
