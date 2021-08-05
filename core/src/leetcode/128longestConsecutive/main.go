package main

import "fmt"

func longestConsecutive(nums []int) int {
	hashMap := make(map[int]struct{})
	for _, v := range nums {
		hashMap[v] = struct{}{}
	}
	res := 0
	for k, _ := range hashMap {
		if _,ok:= hashMap[k-1]; ok {
			 continue
		}
		cur := k
		tempLen := 1
		_, ok2:= hashMap[cur+1]
		for ok2 {
			tempLen ++
			cur ++
			_, ok2 = hashMap[cur+1]
		}
		if tempLen > res {
			res = tempLen
		}
	}
	return res
}

func main() {
	nums := []int {10, 2, 3, 1, 4}
	res := longestConsecutive(nums)
	fmt.Println(res)
}
