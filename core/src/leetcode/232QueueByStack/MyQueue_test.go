package main

import (
	"fmt"
	"testing"

)

var queue = Constructor()

func TestMyQueue_Push(t *testing.T) {
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	fmt.Println(queue.Empty())
}