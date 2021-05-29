package main

func permute(nums []int) [][]int {
	var ans [][]int
	visited := make([]bool, len(nums))

	var backTrace func(curPermute[]int, visited []bool)
	backTrace = func(curPermute []int, visited []bool) {
		if len(curPermute) == len(nums) {

		}

		for i, val := range nums {
			if visited[i] == false {
				visited[i] == true
				backTrace()
			}
		}
	}
}
