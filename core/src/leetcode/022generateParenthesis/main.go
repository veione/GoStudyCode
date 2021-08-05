package main

import "fmt"

func generateParenthesis(n int) []string {
	res := make([]string, 0, 5)
	dfs (&res, "", 0, 0, n)
	return res
}

func dfs(res *[]string, curStr string, left, right, n int){
	if left > n || right > n || right > left{
		return
	}
	if left == right && left == n {
		*res = append(*res, curStr)
	}
	dfs(res, curStr+"(", left+1, right, n)
	dfs(res, curStr+")", left, right+1, n)
}

func main() {
	n := 2
	res := generateParenthesis(n)
	for _, str := range res {
		fmt.Printf("%s ", str)
	}
}