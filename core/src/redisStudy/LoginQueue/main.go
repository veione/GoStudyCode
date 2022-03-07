package main

import "fmt"

func main() {
	for i := 1; i <= 20; i++ {
		waitLevel, waitTime := checkLoginWaitLevel(fmt.Sprintf("player-%d", i))
		fmt.Printf("player-%d: waitLevel:%d  waitTime:%d \n", i, waitLevel, waitTime)
	}
}
