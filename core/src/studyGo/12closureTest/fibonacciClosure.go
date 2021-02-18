package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	f0, f1, f2 := 0, 1, 0
	index := 0
	return func() int{
		if index == 0 {
			index += 1
			return f0
		} else if index == 1 {
			index += 1
			
			return f1
		} else {
			f2 = f0 + f1
			f0 = f1
			f1 = f2
			return f2
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}