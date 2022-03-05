package main

import "fmt"

func killProcess(pid, ppid []int, kill int) []int {
	// 建图
	m := make(map[int][]int, len(pid))

	for i := range pid {
		if ppid[i] != 0 {
			if _, ok := m[ppid[i]]; !ok {
				m[ppid[i]] = make([]int, 0, 1)
			}
			m[ppid[i]] = append(m[ppid[i]], pid[i])
		}
	}
	for key, ints := range m {
		fmt.Print(key, ": ")
		for i := range ints {
			fmt.Print(ints[i], " ")
		}
		fmt.Println()
	}
	res := make([]int, 0, len(pid))
	dfs(kill, m, &res)
	return res
}

func dfs(root int, m map[int][]int, res *[]int) {
	*res = append(*res, root)
	for i := range m[root] {
		dfs(m[root][i], m, res)
	}
}

func main() {
	pid := []int{1, 3, 10, 5}
	ppid := []int{3, 0, 5, 3}

	res := killProcess(pid, ppid, 5)
	for i := range res {
		fmt.Print(res[i], " ")
	}
}
