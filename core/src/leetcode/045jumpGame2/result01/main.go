package main

import (
	"fmt"
	"math"
)

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	minSteps := make([]int, len(nums))
	minSteps[0] = 0
	for k := 1; k < len(minSteps); k++ {
		minSteps[k] = math.MaxInt32
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+nums[i] && j< len(nums); j++ {
			if minSteps[i]+1 < minSteps[j] {
				minSteps[j] = minSteps[i] + 1
			}
		}
	}
	return minSteps[len(minSteps)-1]
}

func main() {
	nums := []int {2,3,0,1,4}
	res := jump(nums)
	fmt.Println(res)
}
