package main

import "fmt"

func Myfunc(nums [][]byte) [][]byte {
	n := len(nums)
	ans := [][]byte{}
	track := []byte{}
	var backtrack func(level int)
	backtrack = func(level int) {
		if level == n {
			ans = append(ans, append([]byte{}, track...))
			return
		}
		for i := 0; i < len(nums[level]); i++ {
			track = append(track, nums[level][i])
			backtrack(level + 1)
			track = track[:len(track)-1]
		}

	}
	backtrack(0)
	return ans
}
func main() {
	ans := Myfunc([][]byte{{'a', 'b', 'c'}, {'1', '2', '3', '4'}, {'A', 'B', 'C'}})
	for i := 0; i < len(ans); i++ {
		fmt.Printf("%v\n", string(ans[i]))
	}
}
