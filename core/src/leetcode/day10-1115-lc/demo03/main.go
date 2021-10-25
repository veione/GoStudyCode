package main

import (
	"fmt"
	"sync"
	"time"
)

var n int = 2
var ch1 ,ch2 chan int = make(chan int), make(chan int)

func foo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ch1
		print("foo")
		ch2 <- 1
	}
}

func bar(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ch2
		print("bar")
		ch1 <- 1
	}
}

func main() {
	ch1 <-1
	var wg sync.WaitGroup
	wg.Add(2)
	go foo(&wg)
	go bar(&wg)

	wg.Wait()
	time.Sleep(time.Second*2)
	fmt.Printf("\n main \n")
}
