package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func foo(n int, ch chan int) {
	defer wg.Done()
	for i := 0; i < 2*n; i++ {
		if i%2 == 0 {
			fmt.Print("foo")
		}
		ch <- i
	}
}
func bar(n int, ch chan int) {
	defer wg.Done()
	for i := 0; i < 2*n; i++ {
		<-ch
		if i%2 == 1 {
			fmt.Print("bar")
		}
	}
}
func main() {
	n := 10
	ch := make(chan int)
	wg.Add(2)
	go foo(n, ch)
	go bar(n, ch)
	wg.Wait()
	fmt.Printf("\n main \n")
}
