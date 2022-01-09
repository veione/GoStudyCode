package main

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text2) == 0 || len(text1) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	res := 0
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j]= dp[i-1][j-1] +1
			}else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
			if dp[i][j] > res {
				res = dp[i][j]
			}
		}
	}
	return res
}

func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}
