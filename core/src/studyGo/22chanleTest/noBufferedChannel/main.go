package main

import (
	"fmt"
	"time"
)

// 会阻塞
func main() {
	ch := make(chan int)

	go func() {
		ch := <-ch
		fmt.Printf("%v \n", ch)
	}()
	time.Sleep(time.Second * 2)
	close(ch)
	time.Sleep(time.Second * 3)
	fmt.Println("world")
}
