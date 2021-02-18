package main

import "fmt"

func fibonacci (n int, c chan int) {
	x, y := 0, 1
	for i :=0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacciRec (x int) {


}

func main() {
	//ch := make(chan int)
	//ch <- 1
	//ch <- 2
	//fmt.Println(<-ch)
	//fmt.Println(<-ch)

	c := make(chan int, 10)
	fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

