package main

import (
	"encoding/json"
	"fmt"
)

type data struct {
	TownID string  `json:"town_id"`
}

func main() {


	dat := data{TownID: "1234"}
	marshal, err:= json.Marshal(dat)
	fmt.Println(err)
	fmt.Println(string(marshal))
}
