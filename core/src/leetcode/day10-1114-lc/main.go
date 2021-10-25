package main

import (
	"fmt"
	"sync"
)

func first(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("first")
	ch <- 1
}

func second(ch1, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch1
	fmt.Println("second")
	ch2 <- 1
}

func third(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	<-ch
	fmt.Println("third")
}
func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(3)
	go first(ch1,&wg)
	go second(ch1, ch2, &wg)
	go third(ch2, &wg)
	wg.Wait()
	fmt.Println("main")
}
