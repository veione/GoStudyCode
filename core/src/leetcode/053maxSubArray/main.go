package main

import "fmt"

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	temp := nums[0]
	for i := 1; i< len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}else {
			nums[i] = 0
		}
	}

}

func main() {
	nums := []int {-2,1,-3,4,-1,2,1,-5}
	res := maxSubArray(nums)
	fmt.Println(res)
}