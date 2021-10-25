package main

import "sync"

func main() {
	lock := sync.Mutex{}
	lock.Unlock()
}
