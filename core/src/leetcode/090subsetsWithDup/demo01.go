package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	var backTrace func(nums, curSlice []int, curIndex int)
	backTrace = func(candidates, curSlice []int, curIndex int) {
		temp := make([]int, len(curSlice))
		copy(temp, curSlice)
		ans = append(ans, temp)
		if curIndex == len(nums) {
			return
		}

		for i := curIndex; i < len(nums); i++ {
			if i > curIndex && nums[i] == nums[i-1] {
				continue
			}
			curSlice = append(curSlice, nums[i])
			backTrace(nums, curSlice, i+1)
			curSlice = curSlice[:len(curSlice)-1]
		}
	}
	sort.Ints(nums)
	backTrace(nums, []int{}, 0)
	return ans
}

func main() {
	nums := []int{1, 2, 2}
	ans := subsetsWithDup(nums)
	for _, an := range ans {
		for _, ints := range an {
			fmt.Printf("%d ", ints)
		}
		fmt.Println()
	}
}
