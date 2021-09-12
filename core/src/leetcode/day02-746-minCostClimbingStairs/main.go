package main

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 1 {
		return cost[0]
	}
	if len(cost) == 2 {
		return min(cost[0], cost[1])
	}
	t1, t2 := cost[0], cost[1]
	for i := 2; i < len(cost); i++ {
		cost[i] = min(t1+cost[i], t2+cost[i])
		t1 = t2
		t2 = cost[i]
	}
	return min(cost[len(cost)-1], cost[len(cost)-2])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
