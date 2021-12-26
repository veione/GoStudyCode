package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return nums[1]
	}
	res1 := _rob(nums[:len(nums)-1])
	res2 := _rob(nums[1:])
	return max(res1, res2)
}



func _rob(nums []int) int {
	first := nums[0]
	second := max(nums[0], nums[1])
	for i := 2; i < len(nums); i ++ {
		temp := second
		second = max(first + nums[i], second)
		first = temp
	}
	return  second
}

func max(a, b int )int {
	if a > b {
		return  a
	}
	return b
}

func main() {
	num := []int {5, 2, 3, 9}
	res := rob(num)
	fmt.Println(res)
}