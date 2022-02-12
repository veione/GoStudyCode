package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)

	tk := time.NewTicker(1 * time.Second)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(debug.Stack())
		}
		tk.Stop()
	}()
	for {
		select {
		case <-tk.C:
			{
				fmt.Println("业务心跳")
			}
		case <-ch1:
			fmt.Println("读取ch1")
		case <-ch2:
			fmt.Println("读取ch2")
		}
	}
}
