package main

import "fmt"

type Student struct {
	Name string
}

func main() {
	list := make(map[string]*Student)

	student := Student{"wtq"}

	list["student"] = &student

	list["student"].Name = "syt"

	fmt.Println(list["student"])
}
