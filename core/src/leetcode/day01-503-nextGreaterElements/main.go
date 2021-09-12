package main

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0, len(nums)*2)
	ans := make([]int, len(nums))
	for i := range ans {
		ans[i] = -1
	}
	for i := 0; i <= len(nums)*2; i++ {
		index := i % len(nums)
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[index] {
			ans[stack[len(stack)-1]] = nums[index]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, index)
	}
	return ans
}
