package main

func rob(nums []int) int {
	dp := make([]int, len(nums))
	if len(nums) > 0 {
		dp[0] = nums[0]
	}
	if len(nums) > 1 {
		dp[1] = max(dp[0], nums[1])
	}
	for i := 2; i < len(nums); i ++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return  dp[len(dp) -1]
}

func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}