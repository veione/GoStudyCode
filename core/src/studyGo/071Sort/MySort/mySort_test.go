package MySort

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	BubbleSort(nums)
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
}

func TestQuickSort(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	QuickSort(nums, 0, len(nums)-1)
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
}

func TestQuickSort2(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	QuickSort2(nums, 0, len(nums)-1)
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
}

func TestMergeSort(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	temps := make([]int, len(nums))
	MergeSort(nums, 0, len(nums)-1, temps)
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}

}

func TestHeapSort(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6, 10, 5}
	HeapSort(nums, len(nums))
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}

}
