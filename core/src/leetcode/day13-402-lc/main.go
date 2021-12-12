package main

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0, len(num))
	for i := range num {
		for k >0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k --
		}
		stack = append(stack, num[i])
	}
	if k >0  {
		stack = stack[:len(stack)-k]
	}
	res := strings.TrimLeft(string(stack), "0")
	if len(res) == 0 {
		return  "0"
	}
	return res
}

func main() {
	res := removeKdigits("10200", 1)
	fmt.Println(res)
}
