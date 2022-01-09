package main

import "time"

type conunter struct {
	windowSize  uint32     // 窗口计数大小
	windowTime  int64      // 窗口计数时间
	curWindow   uint32     // 当前窗口计数
	lastTime    int64      // 上次更新时间
}

func (c *conunter) require() bool {
	nowTime := time.Now().Unix()
	c.lastTime = nowTime

	if nowTime - c.lastTime > c.windowTime {
		c.curWindow = 0
		c.curWindow ++
		return true
	}else {
		if c.curWindow >= c.windowSize {
			return false
		}else {
			c.curWindow ++
			return true
		}
	}
}