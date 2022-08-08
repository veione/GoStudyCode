package main

import "fmt"

func main() {
	str := fmt.Sprintf("%04d", 92236)
	fmt.Printf("str:%v, len:%v", str, len(str))
}
