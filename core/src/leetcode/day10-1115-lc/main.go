package main

import (
	"fmt"
	"sync"
)

var n int = 10

func foo(ch1 <-chan int, ch2 chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ch1
		print("foo")
		ch2 <- 1
	}
}

func bar(ch1 chan<- int, ch2 <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ch2
		print("bar")
		ch1 <- 1
	}
}

func main() {
	ch1, ch2 := make(chan int), make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go foo(ch1, ch2, &wg)
	wg.Add(1)
	go bar(ch1, ch2, &wg)

	ch1 <-1
	wg.Wait()
	fmt.Printf("\n main \n")
}
