package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var ch chan string
var wg sync.WaitGroup

func consumer() {
	go func() {
		defer func() {
			fmt.Println("消费者携程结束")
			wg.Done()
		}()

		for {
			select {
			case rec, ok := <-ch:
				if !ok {
					fmt.Println("消费者 读chan失败")
					return
				}
				fmt.Printf("%s ", rec)
				time.Sleep(200 * time.Millisecond)
			}

		}
	}()
}

func producter() {
	defer func() {
		close(ch)
		fmt.Println("生成者携程结束")
	}()
	for i := 0; i < 50; i++ {
		select {
		case ch <- strconv.Itoa(i):
			//time.Sleep(200 * time.Millisecond)
		default:
			fmt.Printf("chan 已满 写入失败 :%v \n", i)
		}
	}
}

func main() {
	ch = make(chan string, 20)
	wg.Add(1)
	go consumer()

	time.Sleep(2 * time.Second)

	go producter()

	wg.Wait()
	fmt.Println("main 函数结束")
}
