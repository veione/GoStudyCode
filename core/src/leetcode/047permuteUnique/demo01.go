package main

import (
	"fmt"
)

type set map[int]struct{}

func (s set) append(key int) {
	s[key] = struct{}{}
}

func (s set) remove(key int) {
	delete(s, key)
}

func (s set) exit(key int) bool {
	_, ok := s[key]
	return ok
}

func permute(nums []int) [][]int {
	ans := [][]int{}
	visited := make([]bool, len(nums))

	var backTrace func(curPermute []int, visited []bool)
	backTrace = func(curPermute []int, visited []bool) {
		if len(curPermute) == len(nums) {
			temp := make([]int, len(curPermute))
			copy(temp, curPermute)
			ans = append(ans, temp)
			return
		}
		// 使用set 剪枝
		s := set{}
		for i, val := range nums {
			if visited[i] == false && !s.exit(val) {
				visited[i] = true
				curPermute = append(curPermute, val)
				s.append(val)
				backTrace(curPermute, visited)
				visited[i] = false
				curPermute = curPermute[:len(curPermute)-1]
			}
		}
	}
	backTrace(make([]int, 0, len(nums)), visited)
	return ans
}

func main() {
	nums := []int{1, 1, 3}
	ans := permute(nums)
	for _, an := range ans {
		for _, ints := range an {
			fmt.Printf("%d ", ints)
		}
		fmt.Println()
	}
}
