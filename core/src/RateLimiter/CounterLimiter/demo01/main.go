package main

type counter struct {
	windowSize uint32     // 窗口时间内允许的最大访问数量
	curWindowCount uint32 // 当前窗口计数
	lastTime   int64      // 上一次访问时间

}


