package main

import "fmt"

type st struct {
	a int
}

var list []st

func getA()*st {
	return &list[0]
}



func main(){
	/*s := make([]int,0,10)
	fmt.Println( s,len(s),cap(s))
	s = append(s, 1,2,3)
	fmt.Println(s, len(s),cap(s))*/

	//s := &st{
	//	a: 1,
	//}
	//
	//fmt.Println(s)
	//s.a = 2
	//fmt.Println(s)

	for i :=1 ; i<=3; i++ {
		list = append(list, st{a :i })
	}

	fmt.Println(list)
	s := getA()
	s.a = 10
	fmt.Println(list)

}

