package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := make([]int, 0, len(s))
	if len(s) < len(p) {
		return res
	}
	pMap := make(map[byte]int, len(p))
	sMap := make(map[byte]int, len(p))
	for i := range p {
		pMap[p[i]]++
	}
	l, r := 0, 0
	for ; r < len(p); r++ {
		sMap[s[r]]++
	}
	if MapEqual(pMap, sMap) {
		res = append(res, l)
	}
	for l < len(s)-len(p) {
		sMap[s[r]]++
		r++
		sMap[s[l]]--
		l++
		if MapEqual(pMap, sMap) {
			res = append(res, l)
		}
	}
	return res
}

func MapEqual(a, b map[byte]int) bool {
	for key, _ := range a {
		if b[key] != a[key] {
			return false
		}
	}
	return true
}

func main() {
	res := findAnagrams("ababa", "ab")
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}
