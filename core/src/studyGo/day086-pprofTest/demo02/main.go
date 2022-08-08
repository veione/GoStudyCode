package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

func main() {
	http.HandleFunc("/pprof-test", handler)

	fmt.Println("http server start")
	err := http.ListenAndServe("localhost:8090", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(resp http.ResponseWriter, req *http.Request) {
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 0; i < 200; i++ {
		go cyclenum(30000, &wg)
	}

	wg.Wait()

	wb := writeBytes()
	b, err := ioutil.ReadAll(wb)
	if err != nil {
		resp.Write([]byte(err.Error()))
		return
	}
	resp.Write(b)
}

func cyclenum(num int, wg *sync.WaitGroup) {
	slice := make([]int, 0)
	for i := 0; i < num; i++ {
		for j := 0; j < num; j++ {
			j = i + j
			slice = append(slice, j)
		}
	}
	wg.Done()
}

func writeBytes() *bytes.Buffer {
	var buff bytes.Buffer

	for i := 0; i < 30000; i++ {
		buff.Write([]byte{'a' + byte(rand.Intn(10))})
	}
	return &buff
}
