package main

import "fmt"

var (
	a int32  = 1
	b string = "wtq"
	c        = []int{1, 2, 3}
	d        = []string{"wtq", "syt"}
)

func main() {
	fmt.Printf("%T \n", a)
	fmt.Printf("%T:%s \n", b, b)
	fmt.Printf("%T: %v \n", c, c)
	fmt.Printf("%T: %d \n", d, len(c))
}
