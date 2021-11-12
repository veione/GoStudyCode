package main

import "fmt"

func gcd(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func nlcm(nums []int, n int) int {
	if len(nums) == 1 {
		return nums[0]
	} else {
		return lcm(nums[n-1], nlcm(nums, n-1))
	}
}

func main() {
	nums := make([]int, 0, 6)
	for i := 10; i <= 20; i+=2 {
		nums = append(nums, i)
		fmt.Print(i, " ")
	}
	ans := nlcm(nums, 6)
	fmt.Printf("\n ans :%v \n", ans)

}
