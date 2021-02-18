package main

import (
	"fmt"
	"time"
)

func show(t int){
	fmt.Print(t," ")
}

func main()  {

	for i := 1; i< 100; i++ {
		go func() {
			show(i)

		}()
	}
	time.Sleep(10*time.Second)
}
