package main

import "sort"

type event struct {
	uid   int
	tim   int64 //
	event int   // 0: 登录  // 1 登出
}

type events []event

type logs []event

func (l logs) Len() int {
	return len(l)
}

func (l logs) Less(i, j int) bool {
	return l[i].tim < l[j].tim
}

func (l logs) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func maxOnlineCount(userLogs logs) int {
	sort.Sort(userLogs)
	res := 0
	onlineCount := 0
	for _, log := range userLogs {
		if log.event == 0 {
			onlineCount++
			if onlineCount > res {
				res = onlineCount
			}
		} else {
			onlineCount--
		}
	}
	return onlineCount
}

func main() {

}
