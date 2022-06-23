package main

import "fmt"

// 贪心策略 ， 每次选择冷却时间 为0 的 且剩余数量最多的任务
func leastInterval(tasks []byte, n int) int {
	// 冷却时间
	timeMap := make(map[byte]int, 25)
	// 剩余数量
	numMap := make(map[byte]int, 25)
	// 初始化
	for _, char := range tasks {
		timeMap[char] = 0
		numMap[char] += 1
	}
	// 已经执行的任务数量
	count := 0
	// 已经经历的时间
	resCount := 0
	for count < len(tasks) {
		temp := byte('a')
		// 是否有 可选的 任务
		for char, _ := range timeMap {
			if timeMap[char] == 0 && numMap[char] > 0 {
				temp = char
			}
		}
		if temp == 'a' {
			// 没有可选任务，当前轮 待命, 但是要更新 任务的 冷却时间
			for char, _ := range timeMap {
				time := timeMap[char] - 1
				if time < 0 {
					timeMap[char] = 0
				} else {
					timeMap[char] = time
				}
			}
			// 经历时间 增加，轮空 也是要加时间的
			resCount++
			continue
		}

		numMax := 0
		selectChar := byte('b')
		// 贪心的 选择 可选任务中 剩余数量最多的任务
		for char, _ := range timeMap {
			if timeMap[char] == 0 && numMap[char] > numMax {
				numMax = numMap[char]
				selectChar = char
			}
		}
		// 选出的任务 更新 剩余数量 和 冷却时间
		numMap[selectChar] -= 1
		timeMap[selectChar] = n

		// 更新 其他任务的 冷却时间
		for char, _ := range timeMap {
			if char != selectChar {
				time := timeMap[char] - 1
				if time < 0 {
					timeMap[char] = 0
				} else {
					timeMap[char] = time
				}
			}
		}
		// 更新 已执行任务数量 和 已经经历的时间
		resCount++
		count++
	}
	return resCount
}

func main() {
	task := "AAABBB"
	res := leastInterval([]byte(task), 0)
	fmt.Println(res)
}
