package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	num, letter := make(chan bool), make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-num:
				{
					i++
					fmt.Print(i)
					i++
					fmt.Print(i)
					letter <- true
					break
				}
			default:
				print("数字")
				time.Sleep(time.Second*1)
				break
			}
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		word := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		i := 0
		for {
			select {
			case <-letter:
				{
					fmt.Print(string(word[i]))
					i++
					if i >= len(word) {
						return
					}
						fmt.Print(string(word[i]))
					i++
					if i >= len(word) {
						return
					}
					num <- true
					break
				}
			default:
				print("字母")
				time.Sleep(time.Second*1)
				break
			}
		}
	}(&wg)
	num <- true
	strings.Contains()
	wg.Wait()
}
