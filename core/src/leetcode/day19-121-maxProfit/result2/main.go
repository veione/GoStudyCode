package main

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([][2]int, len(prices))
	// 第 i 天不持有 股票
	dp[0][0] = 0
	// 第 i 天持有 股票
	dp[0][1] = -prices[0]

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	return dp[len(dp)-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

