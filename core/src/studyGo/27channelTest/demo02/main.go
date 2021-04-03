package main

import "fmt"

var msg string
var c = make(chan bool)

func hello(){
	msg = "王天奇"
	x := <- c
	fmt.Println(x)
}

func main()  {
	go hello()
	c <- true
	fmt.Println(msg)
}