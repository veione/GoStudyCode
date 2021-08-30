package main

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merge := make([][]int, 0, len(intervals))
	if len(intervals) == 0 {
		return merge
	}
	merge = append(merge, intervals[0])
	for i := 1; i < len(intervals); i++ {
		back := merge[len(merge)-1]
		cur := intervals[i]
		if cur[0] > back[1] {
			merge = append(merge, cur)
		} else {
			back[1] = max(back[1], cur[1])
		}
	}
	return merge
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
