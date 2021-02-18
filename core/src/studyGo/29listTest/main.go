package main

import (
	"container/list"
	"fmt"
)

func main()  {
	ls1 := list.New()

	for i :=1; i< 5; i++ {
		ls1.PushBack(i)
	}
	ls2 := list.New()
	ls2 = ls1
	for em := ls2.Front(); em != nil; em = em.Next(){
		fmt.Print(em.Value, " ")
	}

}
