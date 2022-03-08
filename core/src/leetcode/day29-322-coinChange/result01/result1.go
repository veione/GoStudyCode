package result01

import (
	"math"
	"sort"
)

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	sort.Ints(coins)
	for i := 0; i < len(dp); i++ {
		// 总额为 i ， 最多 由  i张 面额为1的钞票构成
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		for j := range coins {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			} else {
				break
			}
		}
	}
	if dp[len(dp)-1] == math.MaxInt32 {
		return -1
	}
	return dp[len(dp)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
