package main

func findLength(nums1 []int, nums2 []int) int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return 0
	}
	dp := make([][]int, len(nums1))
	for i := range dp {
		dp[i] = make([]int, len(nums2))
	}
	result := 0
	for i := 1; i <= len(nums1); i ++ {
		for j := 1; j <= len(nums2); j++ {
			if nums1[i-1] == nums2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}else {
				dp[i][j] = 0
			}
			if dp[i][j] > result {
				result = dp[i][j]
			}
		}
	}
	return result
}
