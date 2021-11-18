package main

import (
	"io"
)

func main(){
	var w io.Writer
	//w = os.Stdout
	//w = new(bytes.Buffer)
	w.Write(make([]byte, 0, 1))
}

