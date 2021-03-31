package main

import "fmt"

//盛水最多的容器

func maxArea(height []int) int {
	i := 0
	j := len(height)-1
	if len(height) < 2 {
		return 0
	}
	res := 0
	for i<j {
		res = max(res, (j-i)* min(height[i], height[j]))
		if height[i]< height[j]{
			i++
		} else {
			j --
		}
	}
	return res
}

func min(a, b int )int {
	if a < b{
		return  a
	}
	return b
}

func max(a, b int )int {
	if a > b{
		return  a
	}
	return b
}

func main() {
	var heigit = []int {4,3,2,1,4}
	res := maxArea(heigit)
	fmt.Print(res)
}