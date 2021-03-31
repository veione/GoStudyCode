package main

import (
	"encoding/base64"
	"fmt"
)

type info struct {
	ip string
}

func main() {
	info := &info{"127.0.0.1"}
	//msg := "127.0.0.1"
	encoded := base64.StdEncoding.EncodeToString([]byte(info.ip))
	//encoded = base64.StdEncoding.EncodeToString([]byte(encoded))
	fmt.Println(encoded)
	fmt.Println(info.ip)
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	//decoded, err = base64.StdEncoding.DecodeString(string(decoded))
	if err != nil {
		fmt.Println("decode error:", err)
		return
	}
	fmt.Println(string(decoded))
}
