package main

func exist(board [][]byte, word string) bool {
	var backTrace func(visited [][]bool, i, j, index int) bool
	backTrace = func(visited [][]bool, i, j, index int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
			return false
		}
		if visited[i][j] == true {
			return false
		}
		if index == len(word) {
			return true
		}
		flag := false
		if word[index] == board[i][j] {
			visited[i][j] = true
			index++
			flag = backTrace(visited, i, j+1, index) || backTrace(visited, i, j-1, index) ||
				backTrace(visited, i+1, j, index) || backTrace(visited, i-1, j, index)
			if !flag {
				index--
				visited[i][j] = false
			}
		}
		return false
	}
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool,len(board[0]))
	}
	for i:= 0; i < len(board); i++ {
		for j :=0; j< len(board[0]); j++ {
			if backTrace(visited, i, j, 1){
				return true
			}
		}
	}
	return false
}

func main() {
	boards := [][]byte{{'A','B','C','E'},{'S','F','C','S'},{'A','D','E','E'}}
	exist(boards, "AB")
}
