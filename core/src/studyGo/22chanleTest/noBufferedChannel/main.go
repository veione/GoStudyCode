package main

// 会阻塞
func main() {
	ch := make(chan int)
	ch <- 1
	<- ch
}
