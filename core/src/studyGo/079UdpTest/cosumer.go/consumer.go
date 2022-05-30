package main

import (
	"fmt"
	"net"
)

const (
	CHANNEL_SIZE = 10000 // channel 缓冲区

	URL = "msg.4399sy.com:5966"
)

type HwLogConsumer struct {
	ch   chan string  // 数据传输信道
	conn *net.UDPConn //
}

func NewHwLogConsumer() *HwLogConsumer {
	res := &HwLogConsumer{
		ch: make(chan string, CHANNEL_SIZE),
	}
	res.init()
	return res
}

func (consumer *HwLogConsumer) init() {

	remoteAddr, err := net.ResolveUDPAddr("udp", URL)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	consumer.conn, err = net.DialUDP("udp", nil, remoteAddr)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	go func() {
		defer func() {
			close(consumer.ch)
			consumer.conn.Close()
		}()
		for {
			select {
			case data, ok := <-consumer.ch:
				if ok {
					_, err = consumer.conn.Write([]byte(data))
				}
			}
		}
	}()
}
