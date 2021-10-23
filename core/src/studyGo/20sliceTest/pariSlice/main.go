package main

import "fmt"

func f(data [][]string)[][]string{
   data = append(data, []string{"132","123"})
   return data
}


func main() {
   var 	data = [][]string{
      {"1","2"},
   }
   data = f(data)
   fmt.Printf("%v", data)
}
