package main

import (
	"expvar"
	"fmt"
	"time"
)

type handle func()

type Router struct {
	handler map[string]handle
	nameChan chan string
}

func NewRouter() *Router {
	return &Router{
		handler:  make(map[string]handle),
		nameChan: make(chan string, 100),
	}
}

func (rt Router) register(name string, function handle) {
	rt.handler[name] = function
}

func (rt Router) getFunc(name string) func() {
	if f, ok := rt.handler[name]; ok {
		return f
	}
	return nil
}


func(rt Router) loop() {
	tk := time.NewTicker(1*time.Second)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Printf("loop错误: %v", err)
		}
		close(tk.C)
	}()
	for {
		select {
		case name := <-rt.nameChan:
			f := rt.getFunc(name)
			if f != nil {
				f()
			}
		}
	}
}

func print1()  {
	fmt.Println("wtq")
}

func print2()  {
	fmt.Println("123")
}

func main() {
	router := NewRouter()
	router.register("wtq", print1)
	router.register("123", print2)

	go router.loop()


}
