package main

import (
	"fmt"
	"sync"
	"time"
)

func proc() {
	panic("ok")
	//fmt.Println("调用proc")
}

func loop(wg *sync.WaitGroup) {
	tick := time.NewTicker(time.Second * 1)
	defer func() {
		wg.Done()
		if err := recover(); err != nil {
			fmt.Printf("loop错误 %v", err)
			go loop(wg)
		}
		tick.Stop()
	}()
	for {
		select {
		case <-tick.C:
			{
				proc()
			}
		}
	}
}

func
main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go loop(&wg)
	wg.Wait()
	fmt.Errorf("main 结束")
}
