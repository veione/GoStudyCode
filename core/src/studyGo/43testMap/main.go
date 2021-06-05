package main

import "fmt"

type ky struct{
	num int
	id  int 
}

func main()  {
	m := make(map[int][]ky, 0)
	for i:= 1; i< 10; i++ {
		temp := make([]ky, 0, 3)
		for j:=1; j< 5;j ++ {
			temp = append(temp, ky{1,1})
		}
		m[i] = temp
	}
	for key, val := range m {
		fmt.Printf("k:%v, val.len:%v \n", key, len(val))
	
	}
}