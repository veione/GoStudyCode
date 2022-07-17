package main

import (
	"fmt"
	"sort"
)

type digitNum struct {
	digit int
	num   int
}

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}
	digits := make([]digitNum, 0, len(nums))
	for k, v := range m {
		digits = append(digits, digitNum{
			digit: k,
			num:   v,
		})
	}
	sort.Slice(digits, func(i, j int) bool {
		return digits[i].num > digits[j].num
	})
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, digits[i].digit)
	}
	return res
}

func main() {
	var nums = []int{1, 1, 1, 2, 2, 3}
	res := topKFrequent(nums, 2)
	for _, val := range res {
		fmt.Printf("%v ", val)
	}
}
