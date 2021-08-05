package main

import "fmt"

type st struct {
	a int
}

var list []st

func getA()*st {
	return &list[0]
}



func main(){
	array := []int{10, 20, 30, 40}
	slice := make([]int, 6)
	n := copy(slice, array)
	fmt.Println(n,slice)

}

