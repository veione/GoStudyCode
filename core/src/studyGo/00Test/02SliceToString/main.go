package main

import (
	"fmt"
	"strings"
)

func Uint32SliceToString(slice1 []uint32, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice1), " "), delim), "[]")
}

func main() {
	s := []uint32{1,2,3,4,5}
	str := Uint32SliceToString(s, ",")
	fmt.Println(str)
}
