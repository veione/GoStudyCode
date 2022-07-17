package main

import (
	"fmt"
	"strconv"
)

func decodeString(s string) string {
	stack := make([]byte, 0, len(s))
	for _, chr := range s {
		if chr == ']' {
			tempStr := ""
			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if top == '[' {
					sum := ""
					for len(stack) > 0 {
						count := stack[len(stack)-1]
						if count >= '0' && count <= '9' {
							stack = stack[:len(stack)-1]
							sum += string(count)
						} else {
							break
						}
					}
					sum = reverseString(sum)
					res, _ := strconv.Atoi(sum)
					str := reverseString(tempStr)
					resStr := ""
					for i := 1; i <= res; i++ {
						resStr += str
					}
					stack = append(stack, str...)
					break
				} else {
					tempStr += string(top)
				}
			}
		} else {
			stack = append(stack, byte(chr))
		}
	}
	return string(stack)
}

func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func main() {
	res := decodeString("3[a2[c]]")
	fmt.Println(res)
}
