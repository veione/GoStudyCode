package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sub := 999999
	res := 0
	for i, _ := range nums {
		if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		l := i +1
		r := len(nums) -1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(sum - target) < sub {
				sub = abs(sum - target)
				res = sum
			}
			if sum < target {
				l++
			}else if sum > target {
				r--
			}else{
				return sum
			}
		}
	}
	return res
}

func abs(num int)int {
	if num < 0 {
		return -num
	}
	return num
}
