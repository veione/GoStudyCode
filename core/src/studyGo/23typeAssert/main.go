package main

import "fmt"

func fun(arg ...interface{}) {
	fmt.Println(len(arg))


	for i, _ := range arg {
		if val, ok := arg[i].(uint32); !ok{
			fmt.Println("类型错误")
		}else {
			fmt.Println(val)
		}
	}
}

func main()  {
	//s := []int {1,2,3}
	//fmt.Println(s...)
	//fun(s...)
}