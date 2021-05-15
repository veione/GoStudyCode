package main

import (
	"testing"
	"time"
)

// 单独的线程，定时删除 排队队列中过期的成员
func TestClearQueueTimeoutMember(t *testing.T)  {
	tic := time.NewTicker(5*time.Second)
	for {
		select {
		case <-tic.C:
			clearLoginQueueTimeoutMemberTimer()
		default:

		}
	}
}