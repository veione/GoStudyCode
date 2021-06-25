package main

import "fmt"

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ans++
				dfs(grid, i, j)
			}
		}
	}
	return ans
}

func dfs(grid [][]byte, i, j int) {
	grid[i][j] = 0
	if j+1 < len(grid[0]) && grid[i][j+1] == '1' {
		dfs(grid, i, j+1)
	}
	if j-1 >=0 && grid[i][j-1] == '1' {
		dfs(grid, i, j-1)
	}
	if i+1 < len(grid) && grid[i+1][j] == '1' {
		dfs(grid, i+1, j)
	}
	if i-1 >=0 && grid[i-1][j] == '1' {
		dfs(grid, i-1, j)
	}
}


func main() {
	grid := [][]byte{{'1','1','0'},{'1','1','0'},{'1','1','0'}}
	islands := numIslands(grid)
	fmt.Println(islands)
}
