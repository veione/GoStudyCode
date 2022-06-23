package main

func findTargetSumWays(nums []int, target int) int {
	ans := 0
	if len(nums) == 0 {
		return 0
	}
	var backtrace func(index int, sum int)
	backtrace = func(index int, sum int) {
		if index == len(nums) {
			if sum == target {
				ans++
			}
			return
		}
		backtrace(index+1, sum+nums[index])
		backtrace(index+1, sum-nums[index])
	}

	backtrace(0, 0)
	return ans
}

func hammingDistance(x int, y int) int {
	ans := 0
	for s := x ^ y; s > 0; s >>= 1 {
		if s&1 == 1 {
			ans++
		}
	}
	return ans
}
