package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type DiningPhilosophers struct {
	wg                     *sync.WaitGroup
	streamForks            [5]chan interface{}
	missingDoubleForkTimes int
}

func (this *DiningPhilosophers) WantToEat(philosopher int, pickLeftFork func(int), pickRightFork func(int), eat func(int), putLeftFork func(int), putRightFork func(int)) {
	defer this.wg.Done()
	left := (philosopher + 4) % 5
	right := (philosopher + 6) % 5
	for {
		select {
		case this.streamForks[left] <- 1:
			{
				pickLeftFork(philosopher)
				select {
				case this.streamForks[right] <- 1:
					pickRightFork(philosopher)
					eat(philosopher)
					<-this.streamForks[left]
					putLeftFork(philosopher)
					<-this.streamForks[right]
					putRightFork(philosopher)
					fmt.Printf("哲学家%v 完成进餐 \n", philosopher)
					return
				default:
					fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, right)
					<- this.streamForks[left]
					putLeftFork(philosopher)
				}
			}
		default:
			fmt.Printf("Philosopher %d can't pick fork %d.\n", philosopher, left)
			this.missingDoubleForkTimes++
			Think()
		}
	}
}

func Eat(philosopher int) {
	fmt.Printf("===== Philosopher %d have eaten. =====\n", philosopher)
}

func Think() {
	Random := func(max int) int {
		rand.Seed(time.Now().Unix())
		return rand.Int() % (max + 1)
	}
	<-time.After(time.Millisecond * time.Duration(Random(50)))
}

func PickLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, leftNum)
}

func PickRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d picked fork %d.\n", philosopher, rightNum)
}

func PutLeftFork(philosopher int) {
	var leftNum = (philosopher + 4) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, leftNum)

}

func PutRightFork(philosopher int) {
	var rightNum = (philosopher + 6) % 5
	fmt.Printf("Philosopher %d putted fork %d.\n", philosopher, rightNum)

}

func main() {
	dp := DiningPhilosophers{
		wg:                     &sync.WaitGroup{},
		missingDoubleForkTimes: 0,
	}
	for i := range dp.streamForks {
		dp.streamForks[i] =  make(chan interface{}, 1)
	}
	for i := range dp.streamForks {
		dp.wg.Add(1)
		go dp.WantToEat(i, PickLeftFork, PickRightFork, Eat, PutLeftFork, PutRightFork)
	}

	dp.wg.Wait()
	fmt.Println("所有哲学家进餐成功")
}
