package main

import "fmt"

func testFunc(s []int) {
	s = append(s, 1)
	s[1] = 2
}

func main() {
	s := make([]int, 2, 2)
	s[0] = 1
	testFunc(s)
	for i := range s {
		fmt.Print(s[i], " ")
	}
}
