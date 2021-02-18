package main

import (
	"fmt"
	"strconv"
	"strings"
)

// SplitStringToUint64Slice ...
func SplitStringToUint64Slice(src string, sep string) []uint64 {
	strSlice := strings.Split(src, sep)
	var uint64Slice []uint64
	for _, item := range strSlice {
		value, err := strconv.ParseUint(item, 10, 64)
		if err != nil {
			continue
		}
		uint64Slice = append(uint64Slice, value)
	}
	return uint64Slice
}

// SplitStringToUint64Slice2 string转整数切片 example:"[1 2 3 4]" -> []uint32{1, 2, 3, 4}
func SplitStringToUint64Slice2(src string, sep string) []uint64 {
	//去掉字符串左边的"["字符
	str1 := strings.TrimLeft(src, "[")
	//去掉字符串右边的"["字符
	str2 := strings.TrimRight(str1, "]")
	return SplitStringToUint64Slice(str2, sep)
}

func main()  {
	s := []uint64 {1,2,3,4}
	str := fmt.Sprintf("%v",s)
	s1 := SplitStringToUint64Slice2(str, " ")
	fmt.Printf("%v", s1)
}