package main

import (
	"container/list"
	"fmt"
)

func largestRectangleArea(heights []int) int {
	stack := list.New()
	resultLeft := make([]int, len(heights))
	resultRight := make([]int, len(heights))
	for i, val := range heights {
		for stack.Len() > 0 {
			backValue, _ := stack.Back().Value.(int)
			if heights[backValue] >= val{
				stack.Remove(stack.Back())
			}else {
				break
			}
		}
		if stack.Len() >0 {
			backValue, _ := stack.Back().Value.(int)
			resultLeft[i] = backValue
		} else {
			resultLeft[i] = -1
		}
		stack.PushBack(i)
	}
	stack = list.New()
	for i := len(heights)-1; i>=0; i-- {
		for stack.Len() > 0 {
			backValue, _ := stack.Back().Value.(int)
			if heights[backValue] >= heights[i]{
				stack.Remove(stack.Back())
			}else {
				break
			}
		}
		if stack.Len() >0 {
			backValue, _ := stack.Back().Value.(int)
			resultRight[i] = backValue
		} else {
			resultRight[i] = len(heights)
		}
		stack.PushBack(i)
	}

	result := 0
	for i, _ := range heights {
		temp := heights[i] * (resultRight[i] - resultLeft[i] -1 )
		if temp > result {
			result = temp
		}
	}
	return result
}

func maximalRectangle(matrix [][]byte) int {
	res := 0
	heights := make([]int, len(matrix[0]))
	for i, data := range matrix {
		for j, _ := range data {
			if 0 == matrix[i][j] {
				heights[j] = 0
			} else {
				heights[j] ++
			}
		}
		temp := largestRectangleArea(heights)
		if temp > res {
			res = temp
		}
	}
	return res
}

func main() {
	//var raw, column = 4, 5
	//fmt.Scan(&raw, &column)

	var matrix = [][]byte {{1, 0, 1, 0, 0},{1, 0, 1, 1, 1},{1, 1, 1, 1, 1},{1, 0, 0, 1, 0}}

	//matrix := make([4][5]byte, raw)
	//for i := range matrix {
	//	matrix[i] = make([]byte, 0, column)
	//	for j := range matrix[i] {
	//		fmt.Scanf("%d",&matrix[i][j])
	//	}
	//}
	//for i := range matrix {
	//	for j := range matrix[i]{
	//		fmt.Printf(" %v", matrix[i][j])
	//	}
	//	fmt.Println()
	//}

	res := maximalRectangle(matrix)
	fmt.Print(res)
}
