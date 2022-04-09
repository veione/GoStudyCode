package main

import "fmt"

var res [][]int

func arrlyList(nums [][]int) {
	res = make([][]int, 0, len(nums))
	data := make([]int, 0, len(nums))
	dfs(nums, 0, &data)

	for i := range res {
		for _, val := range res[i] {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}

func dfs(nums [][]int, deep int, data *[]int) {
	if deep >= len(nums) {
		temp := make([]int, 0, len(*data))
		for i := 0; i < len(*data); i++ {
			temp = append(temp, (*data)[i])
		}
		res = append(res, temp)
		return
	}
	for _, val := range nums[deep] {
		*data = append(*data, val)
		dfs(nums, deep+1, data)
		*data = (*data)[:len(*data)-1]
	}
}

func main() {
	nums := make([][]int, 0, 5)

	array1 := []int{1, 2, 3}
	array2 := []int{-1, -2, -3, -4}
	array3 := []int{5, 6, 7}
	nums = append(nums, array1)
	nums = append(nums, array2)
	nums = append(nums, array3)
	arrlyList(nums)
}