package main

import "fmt"

func main() {
	m := make(map[int]int)
	for i := 1; i <= 5; i++ {
		m[i] = i
	}
	for k := 0; k < 110; k ++ {
		for _, v := range m {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}

}
