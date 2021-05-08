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

func main() {
	input := []int {2,1,5,6,2,3}
	res := largestRectangleArea(input)
	fmt.Print(res)
}

