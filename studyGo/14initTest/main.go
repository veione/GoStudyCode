package main

import (
	"fmt"
	"studyGo/hello"
)

func init(){
	fmt.Println("1")
}

func init(){
	fmt.Println("2")
}

func init(){
	fmt.Println("3")
}

func main(){
	hello.Hello()
}