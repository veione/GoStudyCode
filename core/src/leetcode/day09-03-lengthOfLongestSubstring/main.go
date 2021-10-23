package main

func lengthOfLongestSubstring(s string) int {
	worldMap := make(map[byte]int)
	right := -1
	res := 0
	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(worldMap, s[i-1])
		}
		for right+1 < len(s) && worldMap[s[right+1]] == 0 {
			worldMap[s[right+1]] += 1
			right++
		}
		res = max(res, right-i+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
