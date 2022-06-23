package day48_406_reconstructQueue

import (
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		if a[0] < b[0] {
			return true
		} else if a[0] == b[0] {
			return a[1] > b[1]
		}
		return false
	})
	ans := make([][]int, len(people))
	for _, person := range people {
		space := person[1] + 1
		for i, an := range ans {
			if an == nil {
				space--
			}
			if space == 0 {
				ans[i] = person
				break
			}
		}
	}
	return ans
}
