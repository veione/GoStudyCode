package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	ans := [][]int{}
	var backTrace func(candidates, curSlice []int, curIndex, curSum int)
	backTrace = func(candidates, curSlice []int, curIndex, curSum int) {
		if curSum == target{
			temp := make([]int, len(curSlice))
			copy(temp, curSlice)
			ans = append(ans, temp)
			return
		}
		if curSum > target{
			return
		}
		for i := curIndex; i < len(candidates); i++ {
			if i > curIndex && candidates[i] == candidates [i -1]{
				continue
			}
			curSlice = append(curSlice, candidates[i])
			backTrace(candidates, curSlice, i+1, curSum+candidates[i])
			curSlice = curSlice[:len(curSlice) -1]
		}
	}
	sort.Ints(candidates)
	backTrace(candidates, []int{}, 0, 0)
	return ans
}


func main() {
	nums := []int{3,5}
	ans := combinationSum2(nums, 8)
	for _, an := range ans {
		for _, ints := range an {
			fmt.Printf("%d ", ints)
		}
		fmt.Println()
	}
}
