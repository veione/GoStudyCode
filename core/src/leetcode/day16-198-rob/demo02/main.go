package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return nums[1]
	}
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i ++ {
		temp := second
		second = max(first + nums[i], second)
		first = temp
	}
	return  second
}

func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}