package main

import "fmt"



// f returns 0
func f1() ( int) {
	var result int = 0
	defer func() {
		result++
	}()
	return result
}

// f returns 1
func f2() (result int) {
	defer func() {
		result++
	}()
	return result
}

func main()  {
	res := f2()
	fmt.Println(res)
}
