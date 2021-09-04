package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int, len(nums1))
	stack := make([]int, 0, len(nums2))

	for _, val := range nums2 {
		for len(stack) > 0 && stack[len(stack)-1] < val {
			numMap[stack[len(stack)-1]] = val
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
		numMap[val] = -1
	}

	res := make([]int, 0, len(nums1))
	for _, val := range nums1 {
		res = append(res, numMap[val])

	}
	return res
}
