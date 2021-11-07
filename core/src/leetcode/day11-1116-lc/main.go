package main

import (
	"fmt"
	"sync"
	"time"
)


var ch1, ch2, ch3 = make(chan int), make(chan int), make(chan int)

func zero(n int, wg *sync.WaitGroup) {
	//defer wg.Done()
	for i := 1; i <= n; i++ {
		<-ch1
		fmt.Print("0")
		if i%2 == 1 {
			ch2 <- 1
		} else {
			ch3 <- 1
		}
	}

}

func oddNum(n int, wg *sync.WaitGroup) {
	//defer wg.Done()
	odd := 1
	for i := 0; i < n; i++ {
		<-ch2
		fmt.Print(odd)
		odd += 2
		ch1 <- 1
	}
}

func even(n int, wg *sync.WaitGroup) {
	//defer wg.Done()
	ev := 2
	for i := 0; i < n; i++ {
		<-ch3
		fmt.Print(ev)
		ev += 2
		ch1 <- 1
	}
}

func main() {
	var n int
	fmt.Scanln(&n)
	wg := sync.WaitGroup{}
	//wg.Add(3)
	go zero(n, &wg)
	go oddNum(n, &wg)
	go even(n, &wg)
	ch1 <- 1
	//wg.Wait()
	time.Sleep(2*time.Second)
	fmt.Printf("\nmain 执行完成 \n")


}
