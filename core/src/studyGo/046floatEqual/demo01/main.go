package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := 3.1432
	b := 3.1456
	// compare a to b
	result := big.NewFloat(a).Cmp(big.NewFloat(b))

	// -1 if a < b
	if result < 0 {
		fmt.Println("a less than b")
	}

	// 0 if a == b
	if result == 0 {
		fmt.Println("a  equals to b")
	}

	// +1 if a > b
	if result > 0 {
		fmt.Println("a 1 more than b")
	}
}
