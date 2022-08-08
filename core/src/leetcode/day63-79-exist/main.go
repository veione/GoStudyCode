package main

import "fmt"

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i, _ := range visited {
		visited[i] = make([]bool, len(board[i]))
	}
	var backTrace func(i, j, index int) bool
	backTrace = func(i, j, index int) bool {
		if j < 0 || j >= len(board[0]) || i < 0 || i >= len(board) {
			return false
		}
		if index < 0 || index >= len(word) {
			return false
		}
		if visited[i][j] == false {
			return false
		}
		if board[i][j] != word[index] {
			return false
		}
		if index == len(word)-1 {
			return true
		}
		visited[i][j] = true
		flag1 := backTrace(i+1, j, index+1)
		flag2 := backTrace(i-1, j, index+1)
		flag3 := backTrace(i, j+1, index+1)
		flag4 := backTrace(i, j-1, index+1)
		visited[i][j] = false
		return flag1 || flag2 || flag3 || flag4
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if backTrace(i, j, 0) {
				return true
			}
		}
	}
	return false
}

func main() {
	var board = [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCD"
	res := exist(board, word)
	fmt.Println(res)
}
