package main

import "math"

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	left := math.MaxInt32
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		left = min(left, prices[i])
		dp[i] = max(dp[i-1], prices[i]-left)
	}
	return dp[len(dp)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
