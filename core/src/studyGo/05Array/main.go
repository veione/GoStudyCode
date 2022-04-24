package main

import "fmt"

type task struct {
	stepList [][]*task // 任务列表
}

func main() {
	var t = [][]*int
	fmt.Printf("%T", t)
}
