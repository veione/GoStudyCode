package main

import "time"

type User struct {
	name string
}

func test() *User{
	a := User{}
	return &a
}

func update() {
	t := time.NewTicker(time.Second*2)
	go func() {
		for range t.C {

		}
	}()
}

func main() {
	//test()
	update()
}
