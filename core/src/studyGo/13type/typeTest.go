package main

import (
	"studyGo/hello"
	"time"
)

// 定义time.Duration的别名为MyDuration
type MyDuration time.Duration

// 为MyDuration添加一个函数
func (m MyDuration) EasySet(a string) {
}

func main() {
	hello.Hello()
}
