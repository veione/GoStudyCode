package main

import "fmt"

func main()  {
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2
	fmt.Println(len(m))
	fmt.Println(m[10])

}
