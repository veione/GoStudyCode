package main

import (
	"fmt"
	"sort"
)

func subsets(nums []int) [][]int {
	ans := make([][]int, 0, 10)
	var backTrace func(path *[]int, curDeep int)
	backTrace = func(path *[]int, curDeep int) {
		temp := make([]int, len(*path))
		copy(temp, *path)
		ans = append(ans, temp)
		if curDeep == len(nums) {
			return
		}
		for i := curDeep; i < len(nums); i++ {
			*path = append(*path, nums[i])
			backTrace(path, i+1)
			*path = (*path)[:len(*path)-1]
		}
	}
	sort.Ints(nums)
	path := make([]int, 0, len(nums))
	backTrace(&path, 0)
	return ans
}

func subsets2(nums []int) [][]int {
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
			curSlice = curSlice[:len(curSlice)-1]
		}
	}
	sort.Ints(nums)
	backTrace(nums, []int{}, 0)
	return ans
}

func main() {
	var nums = []int{1, 2, 3}
	ans := subsets(nums)
	for _, nums := range ans {
		fmt.Printf("%v ", nums)
	}
}
