package main

import "fmt"

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[i] {
				index = j
			}
		}
		if index != i {
			nums[index], nums[i] = nums[i], nums[index]
		}
	}

}

func main() {
	nums := []int{1, 5, 3, 9, 6}
	selectSort(nums)
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
}
