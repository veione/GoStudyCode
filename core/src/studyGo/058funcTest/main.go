package main

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

func (p point) Distance(q point){
	fmt.Println("this is Distance")
}

func (p *point) ScalBy()  {
	fmt.Println("this is ScalBy")
}

const (
	c = 1
)

func fun(c uint32){
	fmt.Printf("%T", c)
}

func main()  {
	////p := point{}
	////q := point{}
	////p.Distance(q)
	////p.ScalBy()
	////
	////
	////t := &point{}
	////t.ScalBy()
	////t.Distance(q)
	//
	////fun(c)
	//s := []int {1,2,3,4}
	//str := fmt.Sprintf("%v",s)
	//str2 :=  strings.Trim(str, "[")
	//str3 :=  strings.Trim(str2, "]")
	//fmt.Println(str3)
	//max := ^uint32(0)
	//fmt.Print(max)

	var t uint32 =5655
	//k := time.NewTicker()
	fmt.Println(strconv.FormatUint(uint64(t), 10))
}
