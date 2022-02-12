package main

import "fmt"

type item struct {
	id  uint32
	num int
}

func main() {
	items := make([]item, 0, 5)
	for i := 0; i < 4; i++ {
		items = append(items, item{uint32(i), i})
	}
	for _, item := range items {
		item.id = 0
	}
	for _, item := range items {
		fmt.Printf("id:%v num:%v \n", item.id, item.num)
	}
}
