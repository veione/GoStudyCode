package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	tMap := make(map[byte]int)
	sMap := make(map[byte]int)
	resLen := math.MaxInt32
	resL := -1
	resR := -1
	for i := 0; i < len(t); i++ {
		tMap[t[i]] += 1
	}
	check := func(sMap, tMap map[byte]int) bool {
		for k, v := range tMap {
			sv, ok := sMap[k]
			if !ok || sv < v {
				return false
			}
		}
		return true
	}
	for l, r := 0, 0; r < len(s); r++ {
		sMap[s[r]] += 1
		for check(sMap, tMap) && l <= r {
			sMap[s[l]] -= 1
			if  r-l+1 < resLen {
				resLen = r - l + 1
				resL = l
				resR = r
			}
			l++
		}
	}
	if resL == -1 {
		return ""
	}
	return s[resL:resR+1]
}

func main() {
	s := "abaacbab"
	t := "abc"
	res:=minWindow(s, t)
	fmt.Println(res)
}