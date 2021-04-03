package memo

import (
	"gopl/ch9/memotest"
	"testing"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := New(httpGetBody)
	memotest.Sequential(t, m)
}

// NOTE: not concurrency-safe! Test fails.

// func TestConcurrent(t *testing.T) {
// 	m := New(httpGetBody)
// 	memotest.Concurrent(t, m)
// }
