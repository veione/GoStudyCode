package main

import "strings"

func reverseWords(s string) string {
	ss := strings.Fields(s)
	reverse(&ss, 0, len(ss) -1)
	strings.Join(ss, " ")
}

func reverse(s *[]string, left, right int) {
	for left < right {
		(*s)[left], (*s)[right] = (*s)[right], (*s)[left]
	}

}
