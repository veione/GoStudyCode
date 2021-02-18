package main

import (
	"encoding/json"
	"fmt"
)

type ItemIdNum struct {
	ItemId uint32 `json:"id"`
	Num    uint32 `json:"num"`
}

func main()  {
	var id = []uint32 {1,2,3,4}
	var num = []uint32 {10,11,12,13}

	items := make([]ItemIdNum, 0, len(id))
	for i, id := range id {
		temp := ItemIdNum{
			ItemId: id,
			Num: num[i],
		}
		items = append(items, temp)
	}
	 marshal, err := json.Marshal(items)
	 if err != nil {
	 	fmt.Println(err)
		 return
	 }
	 fmt.Print(string(marshal))
}
