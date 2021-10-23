package main

import (
	"encoding/json"
	"fmt"
)

type online struct {
	UserID          uint64
	Token           string
	ZoneId          int
	Nick            string
}

func main()  {
	o := online{
		UserID: 001,
		Token:  "dsfsda",
		ZoneId: 0,
		Nick:   "\0001d",
	}
	data, err := json.Marshal(o)
	if err != nil {
		fmt.Printf("错误：%v", err)
	}
	fmt.Println(string(data))
}