package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "123,"
	slice := strings.Split(str, ",")
	for i := range slice {
		fmt.Printf("%v ", slice[i])
	}
}
