package main

import (
	"testing"
	"time"
)

func TestClearQueueTimeoutMember(t *testing.T)  {
	tic := time.NewTicker(5*time.Second)
	for {
		select {
		case <-tic.C:
			clearCrossSrvQueueTimeoutMemberTimer()
		default:

		}
	}
}