package main

import "fmt"

func maxProfit(prices []int) int {
	dp := make([][4]int, len(prices))
	// 完成第一次 买入  buy1
	dp[0][0] = -prices[0]
	// 完成第一次 买卖  sell1
	dp[0][1] = 0
	// 完成第二次买入 buy2
	dp[0][2] = -prices[0]
	// 完成第二次卖出 sell2
	dp[0][3] = 0
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]-prices[i])
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]+prices[i])
	}
	return dp[len(prices)-1][3]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	res := maxProfit(prices)

	fmt.Println(res)
}
