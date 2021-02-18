package main

import "fmt"

var msg string
var c = make(chan bool, 1)

func hello(){
	msg = "王天奇"
	c <- true
}

func main()  {
	go hello()
	x := <- c
	fmt.Println(x)
	fmt.Println(msg)
}