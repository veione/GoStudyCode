package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 划分
func Partition(nums []int, left, right int) int {
	l, r := left, right
	t := l
	for l < r {
		for l < r && nums[r] >= nums[t] {
			r--
		}
		for l < r && nums[l] <= nums[t] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[t] = nums[t], nums[l]
	return l
}

func QuickSort(nums []int, first, end, q int) int {
	if first < end {
		rd := rand.Intn(end-first) + first
		nums[rd], nums[first] = nums[first], nums[rd]
	}

	p := Partition(nums, first, end)
	if p == q {
		return nums[p]
	} else if p > q {
		return QuickSort(nums, first, p-1, q)
	} else {
		return QuickSort(nums, p+1, end, q)
	}
}

func findKthLargest(nums []int, k int) int {
	q := len(nums) - k
	rand.Seed(time.Now().UnixNano())
	return QuickSort(nums, 0, len(nums)-1, q)
}

func main() {
	nums := []int{1, 5, 3, 9, 6, 10, 5}
	res := findKthLargest(nums, 2)
	fmt.Println(res)
}
