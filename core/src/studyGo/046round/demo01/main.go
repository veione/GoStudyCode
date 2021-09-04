package	main

import (
	"fmt"
	"math"
)

func GetRoundInt(value float64) int {
	return int(math.Round(value))
}

func main() {
	nums := []float64 {1.001, 1.2,  0.3, 0.5, 0.51, -0.9}
	for _, num := range nums {
		fmt.Println(GetRoundInt(num))
	}
}
