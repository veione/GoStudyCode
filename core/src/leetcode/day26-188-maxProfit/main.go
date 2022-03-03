package main

import (
	"fmt"
	"math"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)

	if n == 0 || k == 0 {
		return 0
	}
	// n天最多只能完成 n/2次股票买卖
	k = min(n/2, k)

	// buy[i][j] 前i 天 完成 j次 股票买卖 手里还剩余 一只股票时 的 最大利润
	buy := make([][]int, n)
	// sell[i][j] 前i 天 完成 j次 股票买卖 不持有股票 的 最大利润
	sell := make([][]int, n)

	for i := 0; i < n; i++ {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}
	// 边界初始化
	for j := 0; j < k+1; j++ {
		buy[0][j] = math.MinInt32
		sell[0][j] = math.MinInt32
	}
	buy[0][0] = -prices[0]
	sell[0][0] = 0

	for i := 1; i < n; i++ {
		buy[i][0] = max(buy[i-1][0], -prices[i])
		for j := 1; j < k+1; j++ {
			buy[i][j] = max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	res := getMax(sell[n-1])
	return res
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

func getMax(nums []int) int {
	var res int
	for i := range nums {
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	res := maxProfit(2, prices)

	fmt.Println(res)
}
