package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)

	dictMap := make(map[string]bool, len(wordDict))
	for _, val := range wordDict {
		dictMap[val] = true
	}
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			str := s[j:i]
			if dp[j] && dictMap[str] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

func main() {
	s := "abc"
	var dic = []string{"a", "bc"}
	ret := wordBreak(s, dic)
	fmt.Println(ret)
}
