package main

import (
	"fmt"
	"time"
)

var ch1 = make(chan int)
var ch2 = make(chan int)

func first(n int) {
	for i := 0; i < n; i++ {
		<-ch2
		fmt.Println("foo")
		ch1 <- 1
	}
}
func seccond(n int) {
	for i := 0; i < n; i++ {
		<-ch1
		fmt.Println("bar")
		ch2 <- 1
	}
}

func main() {
	n := 3 //n为打印次数
	go first(n)
	go seccond(n)
	ch2 <- 1            //2个线程启动后才开始输出
	time.Sleep(time.Second) //防止主线程退出后，子线程没运行完毕
}
