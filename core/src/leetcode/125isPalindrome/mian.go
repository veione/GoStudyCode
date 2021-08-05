package main

import (
	"fmt"
	"strings"
)

func match(s byte) bool {
	if (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') || (s >= '0' && s <= '9') {
		return true
	}
	return false
}

func isEqual(a, b byte) bool {
	if a != b {
		return (a-'a' == b-'A') || (a-'A' == b-'a')
	}
	return a == b
}

func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		if !match(s[i]) {
			i++
			continue
		}
		if !match(s[j]) {
			j--
			continue
		}
		if strings.ToLower(string(s[i]))==strings.ToLower(string(s[j])) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	str := "OP"
	ret := isPalindrome(str)
	fmt.Println(ret)
}
