package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// 定义三个函数 每个函数 开一个 goroutine 按顺序打印 I love you

func funI(ch1 chan struct{}){
	defer wg.Done()
	for i:= 0; i< 3; i++{
		fmt.Println("I")
	}
	ch1 <- struct{}{}
}

func funLove(ch1, ch2 chan struct{}) {
	defer wg.Done()
	<- ch1
	for i:= 0; i< 3; i++{
		fmt.Println("love")
	}
	ch2 <- struct{}{}
}

func funYou(ch2 chan struct{}) {
	defer wg.Done()
	<-ch2
	for i:= 0; i< 3; i++{
		fmt.Println("you")
	}

}

func main() {
	ch1 := make(chan struct{}, 1)
	ch2 := make(chan struct{}, 1)

	wg.Add(3)
	go funI(ch1)
	go funLove(ch1, ch2)
	go funYou(ch2)
	wg.Wait()
	fmt.Println("结束")
}