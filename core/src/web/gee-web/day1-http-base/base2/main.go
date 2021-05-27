package main

import (
	"fmt"
	"log"
	"net/http"
)

type Engin struct {}

func (engin *Engin) ServeHTTP(w http.ResponseWriter, req *http.Request){
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engin := new(Engin)
	log.Fatal(http.ListenAndServe(":9999", engin))
}