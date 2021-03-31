package main

import "fmt"

func sortColors(nums []int)  {
	n := len(nums)
	p0, p2 := 0, n-1
	for i:=0; i<=p2; i++{
		for i<=p2 && nums[i]==2{
			swap(nums, i, p2)
			p2 --
		}
		if nums[i]== 0{
			swap(nums, i, p0)
			p0 ++
		}
	}
}

func swap(nums []int,  i, j int){
	temp := nums[i]
	nums[i] = nums[j]
	nums[j] = temp
}

func main() {
	var nums = []int {2,0,1}
	sortColors(nums)
	for _, item := range nums {
		fmt.Printf("%d ", item)
	}
}