package main

import (
	"fmt"
	"strconv"
)

type st struct {
	arr []int
	name string
}

func main()  {
	//s := &st{
	//}
	//fmt.Println(len(s.arr))
	var t uint64 = 11223
	fmt.Println(strconv.FormatUint(t, 10))
}
