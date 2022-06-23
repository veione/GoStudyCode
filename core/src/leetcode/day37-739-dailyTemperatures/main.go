package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	length := len(temperatures)
	ans := make([]int, length)
	stack := []int{}
	for i, val := range temperatures {
		for len(stack) > 0 && val > temperatures[stack[len(stack)-1]] {
			preIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[preIndex] = i - preIndex
		}
		stack = append(stack, i)
	}
	return ans
}

func main() {
	t := []int{1, 2, 4, 4, 0}
	res := dailyTemperatures(t)
	fmt.Println(res)
}
