package  main

import "fmt"

func main() {
	for i := 1; ; i++ {
		waitLevel, waitTime := checkLoginWaitLevel(fmt.Sprintf("player-%d", i))
		fmt.Printf("player-%d : %d\n", i, res)
	}
}
