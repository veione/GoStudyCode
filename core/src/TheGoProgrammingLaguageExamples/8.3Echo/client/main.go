package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func mustCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main()  {
	 var wg sync.WaitGroup
	 conn, err := net.Dial("tcp", "localhost:8001")
	 if err != nil {
	 	log.Fatal(err)
	 }
	 defer conn.Close()
	 wg.Add(1)
	 go func() {
	 	wg.Done()
	 	io.Copy(os.Stdout, conn)
	 	log.Println("done")
	 }()
	 mustCopy(conn, os.Stdin)
	 conn.Close()
	 wg.Wait()

}
