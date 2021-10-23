package main

import "fmt"

func main() {
	defer_call()                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        
}

func defer_call(){
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	panic("异常")
}