package main

import "fmt"

// 划分
func Partition(nums []int, first, end int) int {
	i, j := first, end
	for i < j {
		for i < j && nums[i] <= nums[j] {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		for i < j && nums[i] <= nums[j] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return i
}

func QuickSort(nums []int, first, end, q int) int {
	p := Partition(nums, first, end)
	if p == q {
		return nums[p]
	} else if p > q {
		return QuickSort(nums, 0, p-1, q)
	} else {
		return QuickSort(nums, p+1, end, q)
	}
}

func findKthLargest(nums []int, k int) int {
	q := len(nums) - k
	return QuickSort(nums, 0, len(nums)-1, q)
}

func main() {
	nums := []int{1, 5, 3, 9, 6, 10, 5}
	res := findKthLargest(nums, 2)
	fmt.Println(res)
}
