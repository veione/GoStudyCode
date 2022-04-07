package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 1; i < 100; i++ {
		rand.Seed(3)
		fmt.Print(rand.Intn(1000), " ")
	}

}
