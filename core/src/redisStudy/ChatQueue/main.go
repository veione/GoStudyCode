package main

import (
	"fmt"
)

/*
func main() {
	var wg sync.WaitGroup
	for i:=0; i< 4; i++ {
		wg.Add(1)
		go func(k int) {
			defer wg.Done()
			for  {
				j := 1
				res := checkSendCrossSrvChat(fmt.Sprintf("go-%d player-%d",k, j))
				fmt.Printf("go-%d player-%d : %d\n",k, j, res)

			}
		}(i)
	}
	wg.Wait()
}
*/

func main() {
	for i := 1; ; i++ {
		res := checkSendCrossSrvChat(fmt.Sprintf("player-%d", i))
		fmt.Printf("player-%d : %d\n", i, res)
	}
}
