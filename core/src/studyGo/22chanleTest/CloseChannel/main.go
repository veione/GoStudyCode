package main

import (
	"fmt"
	"time"
)

func fibonancci(c, quit chan int)  {
	x, y := 0, 1
	for {
		//t := <- quit
		//fmt.Println(t)
		select {
			case c <- x:
				x, y = y, x+y
			case <- quit:
				fmt.Println("quit")
				return
		}
	}
}

func main()  {
	c, quit := make(chan int), make(chan int)
	go fibonancci(c, quit)

	for i := 0; i< 10; i++ {
		fmt.Print(<-c, " ")
	}
	close(quit)
	time.Sleep(10*time.Second)
}
