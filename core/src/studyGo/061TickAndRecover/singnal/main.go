package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func proc() {
	panic("ok")
	//fmt.Println("调用proc")
}

func loop() {
	tick := time.NewTicker(time.Second * 1)
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("loop错误 %v \n", err)
			go loop()
		}
		tick.Stop()
	}()
	for {
		select {
		case <-tick.C:
			{
				proc()
			}
		}
	}
}

func main() {
	go loop()

	// 信号阻塞
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGPIPE)

	for sig := range ch {
		fmt.Println("[service] 进程收到信号 %s", sig)
		switch sig {
		case syscall.SIGHUP:
		case syscall.SIGPIPE:
		default:
			fmt.Println("[service] 进程收到信号准备退出...")
			close(ch)
			break
		}
	}
	fmt.Errorf("main 结束")
}
