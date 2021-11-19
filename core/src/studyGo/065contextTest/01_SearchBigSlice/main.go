package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 题意：
//假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。限时5秒，使用多个goroutine查找切片中是否存在给定值，在找到目标值或者超时后立刻结束所有goroutine的执行。
//比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，查找的目标值为345，如果切片中存在目标值程序输出:"Found it!"并且立即取消仍在执行查找任务的goroutine。如果在超时时间未找到目标值程序输出:"Timeout! Not Found"，同时立即取消仍在执行查找任务的goroutine。kk

func search(ctx context.Context, wg *sync.WaitGroup, ch chan int, data []int, target int, procIndex int) {
	defer wg.Done()
	for _, value := range data {
		select {
		case <-ctx.Done():
			{
				fmt.Println("线程:%v 被动终止", procIndex)
				return
			}
		default:
		}
		if value == target {
			fmt.Println("线程：%v 找到了***", procIndex)
			ch <- 1
			return
		}
	}
	fmt.Println("线程:%v 正常终止", procIndex)
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 1)
	nums := make([]int, 0, 100000)
	rand.Seed(time.Now().Unix())
	target := 99
	for i := 0; i < 99999; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 99)
	wg := &sync.WaitGroup{}
	for i := 0; i < len(nums); i += 10000 {
		wg.Add(1)
		go search(ctx, wg, ch, nums[i:i+10000], target, i+1)
	}

	select {
	case <-ch:
		{
			fmt.Println("找到了")
			cancel()
		}
	case <-time.After(time.Second * 2):
		{
			fmt.Println("超时")
			cancel()
		}
	}
	wg.Wait()
}
