package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func producter(factor int, p chan<- int){
	for i:= 0; ; i++{
		p<- i*factor
	}
}

func consumer(c <-chan int)  {
	for i := range c {
		fmt.Println(i)
	}
}

func main()  {
	d := make(chan int, 3)
	go producter(3, d)
	go producter(5, d)
	go consumer(d)

	//ctrl+c退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sig)
}