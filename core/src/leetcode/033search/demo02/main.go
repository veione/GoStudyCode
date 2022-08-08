package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			res = mid
			return res
		}
		if nums[left] < nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return res
}

func main() {
	var nums = []int{4, 5, 6, 7, 0, 1, 2}
	index := search(nums, 0)
	fmt.Print(index)
}
