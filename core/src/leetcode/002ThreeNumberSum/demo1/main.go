package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i< len(nums); i++ {
		if i-1 >=0 && nums[i] == nums[i-1]{
			continue
		}
		for j := i+1; j < len(nums); j++ {
			if j >i+1 && j-1 >=0 && nums[j] == nums[j-1]{
				continue
			}
			for k := j+1; k < len(nums); k++ {
				if k > j+1 && k-1 >=0 && nums[k] == nums[k-1]{
					continue
				}
				if 0 == (nums[i] + nums [j] + nums[k] ){
					ans = append(ans, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return ans
}

func main()  {
	var nums = []int {-1,0,1,2,-1,-4}
	ans := threeSum(nums)
	for i:= 0; i < len(ans); i++ {
		for j:= 0; j< len(ans[i]); j++ {
			fmt.Print(ans[i][j], " ")
		}
		fmt.Println()
	}

}
