package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	ans := [][]int{}
	var backTrace func(nums, curSlice []int, curDeep int)
	backTrace = func(candidates, curSlice []int, curDeep int) {
		temp := make([]int, len(curSlice))
		copy(temp, curSlice)
		ans = append(ans, temp)
		if curDeep == len(nums) {
			return
		}

		for i := curDeep; i < len(nums); i++ {
			curSlice = append(curSlice, nums[i])
			backTrace(nums, curSlice, i+1)
			curSlice = curSlice[:len(curSlice) -1]
		}
	}
	sort.Ints(nums)
	backTrace(nums, []int{}, 0)
	return ans
}


func main() {
	nums := []int{1,2,3}
	ans := subsets(nums)
	for _, an := range ans {
		for _, ints := range an {
			fmt.Printf("%d ", ints)
		}
		fmt.Println()
	}
}
