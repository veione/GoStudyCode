package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)


	for {
		select {
		case <-ch1:
			fmt.Println("读取ch1")
		case <-ch2:
			fmt.Println("读取ch2")
		}
		fmt.Println("我最心疼 geigei")
	}
}
