package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var ans [][]int
	sort.Ints(nums)
	for i := 0; i< len(nums); i++ {
		if i-1 >= 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			if 0 == nums[i]+nums[j]+nums[k] {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k-1] == nums[k] {
					k--
				}
				j++
				k--
			} else if 0 > nums[i]+nums[j]+nums[k] {
				j++
			} else {
				k--
			}

		}
	}
	return ans
}

func main()  {
	var nums = []int {-1,0,1,2,-1,-4,-2,-3,3,0,4}
	ans := threeSum(nums)
	for i:= 0; i < len(ans); i++ {
		for j:= 0; j< len(ans[i]); j++ {
			fmt.Print(ans[i][j], " ")
		}
		fmt.Println()
	}

}
