package main

import "fmt"

func countSubstrings(s string) int {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	for l := 2; l <= len(s); l++ {
		for i := 0; i < len(s)-1; i++ {
			j := i + l - 1
			if j > len(s)-1 {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if l <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
		}
	}
	res := 0
	for i := 0; i < len(dp); i++ {
		for j, _ := range dp[i] {
			fmt.Printf("%v ", dp[i][j])
			if dp[i][j] {
				res++
			}
		}
		fmt.Println()
	}
	return res
}

func main() {
	res := countSubstrings("abccba")
	fmt.Println(res)
}
