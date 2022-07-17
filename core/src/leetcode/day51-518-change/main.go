package main

import "fmt"

func change(amount int, coins []int) int {
	ans := 0
	var dfs func(coins []int, index int, sum int)
	dfs = func(coins []int, index int, sum int) {
		if sum > amount {
			return
		} else if sum == amount {
			ans++
		}
		for i := index; i < len(coins); i++ {
			dfs(coins, i, sum+coins[i])
		}
	}
	dfs(coins, 0, 0)
	return ans
}

func main() {
	coins := []int{1, 2}
	ans := change(5, coins)
	fmt.Println(ans)
}
