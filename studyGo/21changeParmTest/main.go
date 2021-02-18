package main

import "fmt"

func c(parm ...interface{}){
	for i, item := range parm {
		if t, ok := item.(uint32); ok{
			fmt.Printf("%v:%v %T\n",i, t, t)
		}
	}
}

func b(parm ...interface{}){
	c(parm...)
}

func main()  {
	var k uint32 =32
	b(k)
	b(uint32(1))
	b(1, 2, 3)
	//a := []interface {"w","wtq"}
	//c(a...)
}
