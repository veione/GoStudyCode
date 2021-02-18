package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)


func main() {
	var wg sync.WaitGroup

	p := NewPublisher(100*time.Millisecond, 10)
	defer p.Close()

	all := p.Subscribe()
	golang := p.SubscribeTopic(func(v interface{}) bool {
		if s, ok := v.(string); ok {
			return strings.Contains(s, "golang")
		}
		return false
	})

	p.Publish("hello,  world!")
	p.Publish("hello, golang!")

	wg.Add(2)
	go func() {
		wg.Done()
		for  msg := range all {
			fmt.Println("all:", msg)
		}

	} ()

	go func() {
		wg.Done()
		for  msg := range golang {
			fmt.Println("golang:", msg)
		}
	} ()

	// 运行一定时间后退出
	//time.Sleep(3 * time.Second)

	wg.Wait()
}
