package main

import "fmt"

func max(a,b int )int {
	if a < b {
		return b
	}
	return a
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func trap(height []int) int {
	length := len(height)

	if length == 0{
		return 0
	}

	leftMax := make([]int, length)
	rightMax := make([]int, length)


	leftMax[0] = height[0]
	for i := 1; i < length; i++{
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[length-1] = height[length -1]
	for i := length -2; i >= 0 ; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	res := 0
	for i, _ := range height {
		res += min(leftMax[i], rightMax[i]) - height[i]
	}
	return res
}

func main() {
	nums := []int {0,1,0,2,1,0,1,3,2,1,2,1}
	res := trap(nums)
	fmt.Println(res)
}
