package main

import (
	"fmt"
	"os"
	"time"
)

func launch()  {
	fmt.Println("发射")
}

func main()  {
	fmt.Println("Commening countdown. press Retrun to abort")

	abort := make(chan  struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for  {
		//fmt.Println(countdown)
		select {
		case <- time.After(10*time.Second):
		case <- abort:
			fmt.Println("Launch aborted!")
			return
		//default:
		//	fmt.Println("默认")
		}
	}
	launch()
}

