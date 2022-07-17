package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &IntHeap{}
	heap.Init(h)
	for _, num := range nums {
		if h.Len() < k {
			heap.Push(h, num)
		} else {
			if num >= (*h)[0] {
				heap.Pop(h)
				heap.Push(h, num)
			}
		}
	}
	return (*h)[0]
}

func main() {
	var arr = []int{3, 4, 1, 4, 5, 3, 10, 0}
	res := findKthLargest(arr, 2)
	fmt.Println(res)
}
