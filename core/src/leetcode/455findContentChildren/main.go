package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	j := 0
	count := 0
	for i := range g {
		for ; j < len(s); j++ {
			if s[j] >= g[i] {
				count++
				j++
				break
			}
		}
		if j == len(s){
			return count
		}
	}
	return count
}

func main() {

}
