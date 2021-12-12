package main

import "fmt"

type kk struct {
	name string
	age  int
}

func main() {
	mp := make(map[int]*kk, 3)
	mp[1] = &kk{}
	mp[2] = &kk{}
	fmt.Println(len(mp))
	delete(mp, 1)
	fmt.Println(len(mp))
}
