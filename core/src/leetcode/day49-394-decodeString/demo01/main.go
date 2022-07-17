package main

import (
	"fmt"
	"strconv"
	"strings"
)

func decodeString(s string) string {
	stack := make([]string, 0, len(s))
	i := 0
	for i < len(s) {
		char := string(s[i])
		if char >= "0" && char <= "9" {
			digit := getDigit(s, &i)
			stack = append(stack, digit)
		} else if (char >= "a" && char <= "z") || (char >= "A" && char <= "Z") || char == "[" {
			stack = append(stack, char)
			i++
		} else {
			tempStr := []string{}
			for stack[len(stack)-1] != "[" {
				tempStr = append(tempStr, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for j := 0; j < len(tempStr)/2; j++ {
				tempStr[j], tempStr[len(tempStr)-j-1] = tempStr[len(tempStr)-j-1], tempStr[j]
			}
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			t := strings.Repeat(getString(tempStr), repTime)
			stack = append(stack, t)
			i++
		}
	}
	return getString(stack)
}

func getDigit(s string, index *int) string {
	digit := ""
	for s[*index] >= '0' && s[*index] <= '9' {
		digit += string(s[*index])
		*index++
	}
	return digit
}

func getString(strs []string) string {
	res := ""
	for i := range strs {
		res += strs[i]
	}
	return res
}

func main() {
	res := decodeString("3[a2[c]]")
	fmt.Println(res)
}
