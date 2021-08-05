package main

import (
	"fmt"
	"io"
)

func max(a , b int) int {
	if a > b{
		return a
	}
	return b
}

func canJump(nums []int) bool {
	right := 0
	for i := 0; i < len(nums); i++ {
		if i <= right{
			right = max(right, i+nums[i])
		}else {
			return false
		}
		if right >= len(nums) -1 {
			return true
		}
	}
	return false
}

func main() {
	for {
		nums := make([]int, 0, 4)
		for {
			var a int
			_, err := fmt.Scan(&a)
			if err == io.EOF{
				break
			}
			nums = append(nums, a)
		}
		res := canJump(nums)
		fmt.Println(res)
	}
}