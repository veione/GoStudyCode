package main

import "fmt"

func main()  {
	natures := make(chan int)
	squares := make(chan int)
	go func() {
		for x:= 0; x<=100; x++ {
			natures <- x
		}
		close(natures)
	}()

	go func() {
		for t := range natures{
			squares <- t*t
		}
		close(squares)
	}()

	for s := range squares{
		fmt.Println(s)
	}

}
