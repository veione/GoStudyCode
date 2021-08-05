package main

import "fmt"

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	for i, item := range dp {
		for j, _ := range item {
			dp[i][j] = true
		}
	}
	max_len := 1
	begin := 0
	// lenth 字符串长度
	for lenth := 2 ; lenth<= len(s); lenth++ {
		// i 左边界
		for i:=0; i< len(s); i++ {
			// j 右边界
			j := i + lenth -1
			if j >= len(s){
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			}else {
				if j -i < 3 {
					dp[i][j] = true
				}else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && j - i +1 > max_len{
				max_len = j - i + 1
				begin = i
			}
		}
	}
	return s[begin:max_len+begin]
}

func main() {
	str := "babad"
	res := longestPalindrome(str)
	fmt.Println(res)
}