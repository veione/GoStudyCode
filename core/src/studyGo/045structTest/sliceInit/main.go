package main

import "fmt"

type st struct {
	name string
	like []int
}

func main() {
	s := &st{
		name: "wtq",
	}
	fmt.Printf("len:%v cap:%v", len(s.like), cap(s.like))
	for i := 0; i< 3; i++ {
		s.like = append(s.like, i)
	}
	fmt.Printf("%s  %v",s.name, s.like)
}