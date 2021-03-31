package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(0)

	if r.Len() > 0 {
		r.Value = 1
		r = r.Next()
	}
	r.Do(func(i interface{}) {
		k, ok := i.(int)
		if ok {
			fmt.Print(k)
		}
	})
}

