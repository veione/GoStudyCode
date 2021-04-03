package main

import "fmt"

func addData(dataMap map[string]string , IP string) {
	dataMap["ip"] = IP

}


func main()  {
	m := make(map[string]string)
	addData(m, "10.10.1.1")
	for s, s2 := range m {
		fmt.Println(s,s2)
	}
}