package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	stack := make([]int, 0, len(temperatures))
	ans := make([]int, len(temperatures))
	for i, value := range temperatures {
		if len(stack) > 0 {
			for len(stack) > 0 && temperatures[stack[len(stack)-1]] < value {
				ans[stack[len(stack)-1]] = i - stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		}
		stack = append(stack, i)
		ans[i] = 0
	}
	return ans
}

func main()  {
	nums := []int {73,74,75,71,69,72,76,73}
	ans := dailyTemperatures(nums)
	for _, an := range ans {
		fmt.Print(an, " ")
	}
}
