package main

import (
	"fmt"
	"io"
)

func main() {
	consumer := NewHwLogConsumer()
	var str string
	for {
		_, err := fmt.Scan(&str)
		consumer.ch <- str
		if err == io.EOF {
			break
		}
	}
	select {}
}
