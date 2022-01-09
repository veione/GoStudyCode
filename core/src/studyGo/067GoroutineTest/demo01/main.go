package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {

		}
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("helow world")
}
