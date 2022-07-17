package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

var raddr = flag.String("raddr", "10.1.1.74:10000", "remote server address")
var msgChan chan string
var conn *net.UDPConn
var wg sync.WaitGroup

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
}

func sendMsgLoop(wg *sync.WaitGroup) {
	for {
		select {
		case data, ok := <-msgChan:
			{
				if !ok {
					return
				}
				fmt.Printf("读取输入 :%v \n", data)
				// write a message to server
				_, err := conn.Write([]byte(data))
				if err != nil {
					log.Println(err)
				} else {
					fmt.Println("Packet sent to: ", *raddr)
				}

				// Receive response from server
				buf := make([]byte, 1024)
				rn, rmAddr, err := conn.ReadFromUDP(buf)
				if err != nil {
					log.Println(err)
				} else {
					fmt.Printf("received from: %v, data: %s\n", rmAddr, string(buf[:rn]))
				}
			}
		}
	}
}

func inputMsg(wg *sync.WaitGroup) {
	wg.Done()
	var str string
	for {
		_, err := fmt.Scan(&str)
		if err == io.EOF {
			break
		}
		select {
		case msgChan <- str:
		default:
			fmt.Println("通道已满")
		}
	}
}

func main() {
	// Resolving Address
	remoteAddr, err := net.ResolveUDPAddr("udp", *raddr)
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	msgChan = make(chan string, 1024)

	// Make a connection
	//tmpAddr := &net.UDPAddr{
	//	IP:   net.ParseIP("10.1.1.74"),
	//	Port: ,
	//}

	conn, err = net.DialUDP("udp", nil, remoteAddr)
	// Exit if some error occured
	if err != nil {
		log.Fatalln("Error: ", err)
	}

	wg.Add(1)
	go inputMsg(&wg)
	wg.Add(1)
	go sendMsgLoop(&wg)

	wg.Wait()

	conn.Close()
}
