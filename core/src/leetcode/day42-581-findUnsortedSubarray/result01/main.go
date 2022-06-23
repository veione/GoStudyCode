package main

import (
	"sort"
)

func findUnsortedSubarray(nums []int) int {
	copyNums := make([]int, len(nums))
	copy(copyNums, nums)
	sort.Ints(copyNums)
	left := 0
	right := len(nums) - 1
	for left < len(nums) {
		if copyNums[left] == nums[left] {
			left++
		} else {
			break
		}
	}
	for right >= 0 {
		if copyNums[right] == nums[right] {
			right--
		} else {
			break
		}
	}
	if right > left {
		return right - left + 1
	}
	return 0
}

func main() {

}
