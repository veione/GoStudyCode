package main

import (
	"fmt"
	"sync"
)

func print(i int, wg *sync.WaitGroup) {
	fmt.Printf("%d ", i)

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, _ := range a {
		wg.Add(1)
		go print(i, &wg)
	}
	wg.Wait()
}
