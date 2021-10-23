package main

import (
	"fmt"
	"net"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("panic , err \n")
		}
	}()
	tcpAddr, err := net.ResolveTCPAddr("tcp4", "10.1.1.5")

	if err != nil {
		fmt.Printf("[net] addr resolve error addr:%v, err:%v \n", tcpAddr, err)
	}

	ln, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Printf("%v", err)
	}
	var tempDelay time.Duration
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if max := 1 * time.Second; tempDelay > max {
					tempDelay = max
				}
				fmt.Printf("accept error: %v; retrying in %v \n", err, tempDelay)
				time.Sleep(tempDelay)
				continue
			}
			return
		}
		tempDelay = 0
		// Try to open keepalive for tcp.
		conn.SetKeepAlive(true)
		conn.SetKeepAlivePeriod(1 * time.Minute)
		// disable Nagle's algorithm.
		conn.SetNoDelay(true)
		go func() {

		}()
	}
}
