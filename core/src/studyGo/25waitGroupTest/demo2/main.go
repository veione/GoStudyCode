package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("golang1")
		wg.Done()
	}()
	go func() {
		fmt.Println("golang2")
		wg.Done()
	}()
	wg.Wait()
}

