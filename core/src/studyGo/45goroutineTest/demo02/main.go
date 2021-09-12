package main

import "fmt"

var limit = make(chan int, 4)

func main() {
	works := []int {4,3,2,1}
	for _, w := range  works{
		go func(w int) {
			limit <- 1
			fmt.Println(w)
			<-limit
		}(w)
	}
	select{}
}