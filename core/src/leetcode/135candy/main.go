package main

import "fmt"

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		if i > 0 && ratings[i] > ratings[i-1] {
			candys[i] = candys[i-1] + 1
		} else {
			candys[i] = 1
		}
	}
	res := 0
	right := 0
	for i := len(ratings) - 1; i >= 0; i-- {
		if i < len(ratings)-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		res += max(candys[i], right)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int {1,0,2}
	res := candy(nums)
	fmt.Print(res)
}