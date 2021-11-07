package main

import (
	"fmt"
	"time"
)

func main()  {
	now := "2006-01-02 00:00:00"
	t, err1 := time.Parse("2006-01-02 00:00:00", now)
	if err1 == nil{
		fmt.Println(err1.Error())
	}
	fmt.Printf("hour：%v  minute:%v  second:%v", t.Hour(), t.Minute(), t.Second())
}
