package main

import "fmt"

func searchRange(nums []int, target int) []int {
	left := binarySearchFirst(nums, target)
	right := binarySearchLatest(nums, target)
	if right == -1 || left == -1 {
		return []int{-1, -1}
	}
	return []int {left, right}
}

//  查找 第一个 = target的 数据
func binarySearchFirst(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target{
			right = mid -1
		} else if nums[mid] < target {
			left = left + 1
		}else {
			index = mid
			right =mid -1
		}
	}
	return index
}
//  查找最后一个 = target 的数据
func binarySearchLatest(nums []int, target int) int {
	left, right := 0, len(nums)-1
	index := -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target{
			right = mid - 1
		} else if nums[mid] < target {
			left = left + 1
		}else {
			left = left + 1
			index = mid
		}

	}
	return index
}
func main() {
	nums := []int{5,7,7,8,8,10}
	target := 6
	ans := searchRange(nums, target)
	for _, an := range ans {
		fmt.Print(an, " ")
	}
}
