package main

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return x*sy+y > y*sx+x
	})

	if nums[0] == 0 {
		return "0"
	}
	res := make([]byte, 0, len(nums))
	for i := range nums {
		res = append(res, strconv.Itoa(nums[i])...)
	}
	return string(res)
}
