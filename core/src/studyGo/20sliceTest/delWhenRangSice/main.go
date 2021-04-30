package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range slice {
		if value%3 == 0 { // remove 3, 6, 9
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	fmt.Printf("%vn", slice)
}