package dfs

func wordBreak(s string, wordDict []string) bool {
	dictMap := make(map[string]struct{}, len(wordDict))
	cache := make(map[int]bool, len(wordDict))
	for _, val := range wordDict {
		dictMap[val] = struct{}{}
	}
	var dfs func(index int) bool
	dfs = func(index int) bool {
		if index == len(s) {
			return true
		}
		if val, ok := cache[index]; ok {
			return val
		}
		for i := index + 1; i <= len(s); i++ {
			str := s[index:i]
			if _, ok := dictMap[str]; ok && dfs(i) {
				cache[i] = true
				return true
			}
		}
		cache[index] = false
		return false
	}
	res := dfs(0)
	return res
}
