package main

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	sum := nums[0]
	res := math.MaxInt32
	for right < len(nums) && left <= right {
		if sum < target {
			right ++
			if right == len(nums) {
				break
			}
			sum += nums[right]
		} else {
			if (right - left + 1) < res{
				res = right - left + 1
			}
			sum -= nums[left]
			left ++
		}
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func main() {
	nums :=[]int {1,4,4}
	target := 4
	res := minSubArrayLen (target, nums)

	fmt.Println(res)
}
