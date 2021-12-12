package main

var c = make(chan int, 1)
var a string

func f() {
	a = "hello, world"
	c <- 0
}
func main() {
	go f()
	<-c
	print(a)
}
