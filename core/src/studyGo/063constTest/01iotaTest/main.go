package main

import "fmt"

const (
	a = iota
	b = iota
)
const (
	name = "menglu"
	c    = "dfasdf"
	d    = iota
	e    = "dsfj"
	f    = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(f)

}
