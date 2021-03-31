package main

import "fmt"

func nextPermutation(nums []int)  {
	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i -1]{
		i --
	}
	if i == 0{
		reverse(nums, 0, len(nums)-1)
		return
	}
	i --
	j := len(nums) -1
	for j > i && nums[j] <= nums[i]{
		j --
	}
	swap(nums, i, j)
	reverse(nums, i+1, len(nums)-1)
}

func swap (nums []int, i, j int){
	temp := nums[i]
	nums[i] = nums[j]
	nums[j]	= temp
}

func reverse(nums []int, i, j int){
	for i<j{
		swap(nums, i, j)
		i++
		j--
	}
}

func main() {
	var nums = []int {1,5,1}
	nextPermutation(nums)
	for _, item := range nums {
		fmt.Print(item," ")
	}
}


