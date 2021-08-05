package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func wordCount(strs []string, grCount int) map[rune]int {
	ans := make(map[rune]int)
	ansChan := make(chan map[rune]int, grCount)
	l := len(strs) / grCount
	if l == 0{
		grCount = 1
	}

	for i:=0; i< grCount; i++ {
		wg.Add(1)
		start := i*l
		end := start + l
		if i == grCount -1 {
			end = len(strs)
		}
		go worker(strs[start:end], ansChan)
	}
	wg.Wait()
	close(ansChan)

	//for{
	//
	//	select{
	//	case charMap := <- ansChan:
	//		for k, v := range charMap {
	//			ans[k] += v
	//		}
	//	}
	//}
	for charMap := range ansChan {
		for k, v := range charMap {
			ans[k] += v
		}
	}
	return ans
}

func worker(strs []string, ansChan chan<- map[rune]int) {
	defer wg.Done()
	temp := make(map[rune]int)
	for _, str := range strs {
		for _, ch := range str {
			temp[ch] ++
		}
	}
	ansChan <- temp
}

func main()  {
	strs := []string{"abc","abc","ab","a"}
	mapCount := wordCount(strs, 3)
	for k, v := range mapCount {
		fmt.Printf("k:%v  v:%d \n", k, v)
	}

}