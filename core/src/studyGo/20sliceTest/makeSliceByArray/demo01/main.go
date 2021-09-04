package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	df s slice1 := array[2:5:6] df
	fmt.Printf("slice1:%p len:%d cap:%d %v \n", &slice1, len(slice1), cap(slice1), slice1)
	slice2 := array[1:3:6]
	fmt.Printf("slice2:%p len:%d cap:%d %v \n", &slice2, len(slice2), cap(slice2), slice2)
}
