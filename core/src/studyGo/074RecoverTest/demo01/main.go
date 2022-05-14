package main

import "fmt"

func loop() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic")
			go loop()
		}
	}()
	for i := 10; i >= -10; i-- {
		fmt.Print(i/i, " ")
	}
	fmt.Println("结束")
}

func main() {
	go loop()
	select {}
}
