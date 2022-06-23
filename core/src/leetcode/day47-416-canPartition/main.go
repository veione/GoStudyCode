package main

import (
	"fmt"
	"sort"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	target := sum / 2
	res := false
	var backTrace func(nums []int, curIndex, curSum int) bool
	backTrace = func(nums []int, curIndex, curSum int) bool {
		if curSum == target {
			res = true
			return true
		}
		if curSum > target {
			return false
		}
		for i := curIndex; i < len(nums); i++ {
			r |= backTrace(nums, i+1, curSum+nums[i])
			if r {
				break
			}
		}
		return false
	}
	sort.Ints(nums)
	backTrace(nums, 0, 0)
	return res
}

func main() {
	res := canPartition([]int{1, 5, 11, 5})
	fmt.Println(res)
}
