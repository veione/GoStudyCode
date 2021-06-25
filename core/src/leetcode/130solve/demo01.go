package main

func solve(board [][]byte)  {
	flag := make([][]bool, len(board))
	for i := range flag {
		flag[i] = make([]bool, len(board[0]))
	}

	for i := range board {
		dfs(board, flag, i, 0)
		dfs(board, flag, i, len(board[0]) -1)
	}

	for j := range board[0] {
		dfs(board,flag, 0, j)
		dfs(board,flag, len(board)-1, j)
	}

	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' && !flag[i][j]{
				board[i][j] = 'X'
			}
		}
	}
}

func dfs(board [][]byte, flag [][]bool, i, j int){
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return
	}
	if flag[i][j] || board[i][j] != 'O' {
		return
	}
	flag[i][j] = true

	dfs(board, flag, i+1, j)
	dfs(board, flag, i-1, j)
	dfs(board, flag, i, j+1)
	dfs(board, flag, i, j-1)
}