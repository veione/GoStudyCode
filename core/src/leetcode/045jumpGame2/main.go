package main

import "fmt"

func jump(nums []int) int {
	maxIndex := 0
	ans := 0
	for maxIndex < len(nums)-1{
		tempIndex := maxIndex
		index := maxIndex
		for i := 1; maxIndex+i <len(nums) && i<= nums[maxIndex]; i++{
			if maxIndex + i + nums[maxIndex+i]>tempIndex{
				tempIndex = maxIndex + i + nums[maxIndex+i]
				index = maxIndex +i
			}
		}
		maxIndex = index
		ans ++
	}
	return ans
}

func main() {
	var nums = []int {2,1}
	ans := jump(nums)
	fmt.Print(ans)
}