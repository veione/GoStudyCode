package main

func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))

	dp[0][0] = -prices[0] // 第i 天  持有股票 时的 最大利润
	dp[0][1] = 0          // 第i 天  不持有股票 且 处于冷冻期  时的最大利润
	dp[0][2] = 0          // 第i 天  不持有股票 且 不处与冷冻期  时的 最大利润

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = max(dp[i-1][2], dp[i-1][1])
	}
	n := len(prices) - 1
	return max(dp[n][1], dp[n][2])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
