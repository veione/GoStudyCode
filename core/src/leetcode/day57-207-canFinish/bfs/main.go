package bfs

func canFinish(numCourses int, prerequisites [][]int) bool {
	count := 0
	queue := make([]int, 0, numCourses)
	inCount := make([]int, numCourses)
	for _, edge := range prerequisites {
		inCount[edge[0]]++
	}
	for i, val := range inCount {
		if val == 0 {
			queue = append(queue, i)
		}
	}

	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		count++
		for _, edge := range prerequisites {
			if edge[1] == start {
				inCount[edge[0]]--
				if inCount[edge[0]] == 0 {
					queue = append(queue, edge[0])
				}
			}
		}
	}
	return count == numCourses
}
