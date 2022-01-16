package main

import (
	"fmt"
	"sort"
)

type AreaSatisfaction struct {
	AreaID  uint32 // 区域ID
	Prosper int32  //
}

func main() {
	aresDatas := make([]*AreaSatisfaction, 0, 5)
	for i := 1; i < 6; i++ {
		aresDatas = append(aresDatas, &AreaSatisfaction{AreaID: uint32(i), Prosper: int32(i)})
	}
	sort.Slice(aresDatas, func(i, j int) bool {
		return aresDatas[i].Prosper > aresDatas[j].Prosper
	})
	for i, _ := range aresDatas {
		fmt.Println(aresDatas[i].AreaID, "", aresDatas[i].Prosper)
	}
}
