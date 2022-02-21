package MySort

import (
	"fmt"
)

func BubbleSort(nums []int) {
	n := len(nums)
	for i := 1; i <= n-1; i++ {
		for j := 0; j <= n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func SelectSort(nums []int) {
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

// 划分1 默认是 以最左边的元素 为基准进行划分
func Partition(nums []int, first, end int) int {
	i, j := first, end
	for i < j {
		for i < j && nums[i] <= nums[j] {
			j--
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		for i < j && nums[i] <= nums[j] {
			i++
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		}
	}
	return i
}

func Partition2(nums []int, left, right int) int {
	l, r := left, right
	t := l
	for l < r {
		for l < r && nums[r] >= nums[t] {
			r--
		}
		for l < r && nums[l] <= nums[t] {
			l++
		}
		if l < r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	nums[l], nums[t] = nums[t], nums[l]
	return l
}

//sentry := rand.Intn(right-left) + left
func QuickSort(nums []int, first, end int) {
	if first < end {
		p := Partition(nums, first, end)
		QuickSort(nums, 0, p-1)
		QuickSort(nums, p+1, end)
	}
}

func QuickSort2(nums []int, first, end int) {
	if first < end {
		p := Partition2(nums, first, end)
		QuickSort2(nums, 0, p-1)
		QuickSort2(nums, p+1, end)
	}
}

func merge(nums []int, left, mid, right int, temps []int) {
	i, j, k := left, mid+1, 0
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			temps[k] = nums[i]
			k++
			i++
		} else {
			temps[k] = nums[j]
			k++
			j++
		}
	}
	for i <= mid {
		temps[k] = nums[i]
		i++
		k++
	}
	for j <= right {
		temps[k] = nums[j]
		j++
		k++
	}
	k = 0
	for t := left; t <= right; t++ {
		nums[t] = temps[k]
		k++
	}
}

func MergeSort(nums []int, left, right int, temps []int) {
	if left < right {
		mid := (right + left) / 2
		MergeSort(nums, left, mid-1, temps)
		MergeSort(nums, mid+1, right, temps)
		merge(nums, left, mid, right, temps)
	}
}

// 向下调整堆
// start: 需要调整的堆， 堆顶    end: 需要调整的堆 的最后一个叶子节点
func ShiftHeap(nums []int, start, end int) {
	i := start
	j := 2*i + 1 // i的左孩子
	for j < end {
		if j < end-1 && nums[j] < nums[j+1] {
			j++
		}
		if nums[i] > nums[j] {
			break
		} else {
			nums[i], nums[j] = nums[j], nums[i]
			i = j
			j = 2*i + 1
		}
	}
}

func HeapSort(nums []int, n int) {
	// 自底向上 建堆
	for i := (n - 1) / 2; i >= 0; i-- {
		ShiftHeap(nums, i, n)
	}
	for i := range nums {
		fmt.Printf("%d ", nums[i])
	}
	fmt.Printf("\n")
	for i := 1; i <= n-1; i++ {
		nums[0], nums[n-i] = nums[n-i], nums[0]
		ShiftHeap(nums, 0, n-i)
	}
}
