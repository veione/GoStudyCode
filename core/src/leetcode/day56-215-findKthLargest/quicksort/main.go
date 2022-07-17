package main

import "fmt"

func findKthLargest(nums []int, k int) int {
	target := len(nums) - k
	var quickSort func([]int, int, int) int
	quickSort = func(nums []int, left, right int) int {
		p := partition(nums, left, right)
		if p == target {
			return nums[p]
		} else if p > target {
			return quickSort(nums, left, p-1)
		} else {
			return quickSort(nums, p+1, right)
		}
	}
	res := quickSort(nums, 0, len(nums)-1)
	return res
}

func partition(nums []int, left, right int) int {
	l, r := left, right
	temp := left
	for l < r {
		for l < r && nums[r] >= nums[temp] {
			r--
		}
		for l < r && nums[l] <= nums[temp] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}

	}
	nums[l], nums[temp] = nums[temp], nums[l]
	return l
}

func main() {
	var arr = []int{3, 4, 1, 4, 5, 3, 10, 0}
	res := findKthLargest(arr, 2)
	fmt.Println(res)
}
