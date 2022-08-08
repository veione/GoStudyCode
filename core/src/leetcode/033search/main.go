package main

import "fmt"

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		//左边部分是有序数组
		if nums[l] <= nums[m] {
			if nums[l] <= target && nums[m] > target {
				r = m - 1
			} else {
				l = m + 1
			}
		} else {
			// 右部有序
			if nums[m] < target && nums[r] >= target {
				l = m + 1
			} else {
				r = m - 1
			}
		}
	}
	return -1
}

func main() {
	var nums = []int{3, 4, 5, 6, 7, 8, 1, 2}
	index := search(nums, 2)
	fmt.Print(index)
}
