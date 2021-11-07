package main

import (
	"fmt"
	"sync"
	"time"
)

const LEN = 2

type mylock struct {
	sync.Mutex
}

var (
	lock mylock
	set = make(map[string]int)
)

func hydrogen() {
	lock.Lock()
	defer lock.Unlock()
	set["H"]++
	fmt.Println(">>> H", set)
}

func oxygen() {
	lock.Lock()
	defer lock.Unlock()
	set["O"]++
	fmt.Println(">>> O", set)
}

func synthesizeH2O() {
	for flag := false; !flag; {
		lock.Lock()
		if set["H"] >= 2 && set["O"] >= 1 {
			set["H"], set["O"] = set["H"]-2, set["O"]-1
			fmt.Println(">>> H2O", set)
			flag = true
		}
		lock.Unlock()
	}

}

func main() {
	for i := 0; i < LEN; i++ {
		go hydrogen()
		go hydrogen()
		go oxygen()
		go synthesizeH2O()
	}
	time.Sleep(time.Duration(1000) * time.Millisecond)
}

