package reverseSort

import (
	"fmt"
	"sort"
	"testing"
)

// 降序排序
func TestReverseSort(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
}

func TestSortSearch(t *testing.T) {
	nums := []int{1, 5, 3, 9, 6}
	sort.Ints(nums)
	index := sort.SearchInts(nums, 5)
	fmt.Printf("%d ", index)
}
