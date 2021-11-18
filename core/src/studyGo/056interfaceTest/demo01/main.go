package main

import "fmt"

//1
type I interface {
	Get() int
	Set(int)
}

//2
type S struct {
	Age int
}

func(s S) Get()int {
	return s.Age
}

func(s *S) Set(age int) {
	s.Age = age
}

//2
type K struct {
	Age int
}

func(k K) Get()int {
	return k.Age
}

func(k K) Set(age int) {
	k.Age = age
}

//3
func f(i I){
	i.Set(10)
	fmt.Println(i.Get())
}

func main() {
	s := S{}
	var i I
	i = &s
	switch t := i.(type) {
	case *S:
		{
			fmt.Println("是S类型", t)
		}
	case *K:
		{
			fmt.Println("是K类型", t)
		}
		
	}
}
