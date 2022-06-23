package result

func lengthOfLongestSubstring(s string) int {
	worldMap := make(map[byte]int)
	right := -1
	res := 0
	for i := 0; i < len(s);  {
		if i > 0 {
			delete(worldMap, s[i-1])
		}
		for right+1 < len(s) && worldMap[s[right+1]] == -1 {
			worldMap[s[right+1]] = right+1
			right++
		}
		res = max(res, right-i+1)
		if right < len(s) {
			i = 
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
