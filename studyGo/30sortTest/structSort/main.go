package main

import (
	"fmt"
	"sort"
)


func main() {
	var stuGrade grades

	for i:=1; i <=5; i++ {
		var en, ch, ma uint32
		fmt.Scanf("%d,%d,%d", &ch, &en, &ma)
		fmt.Printf("ch:%d, en:%d, ma:%d\n", ch, en, ma)
		stuGrade = append(stuGrade, &grade{chinese: ch, math: ma, english: en})
	}
	sort.Sort(stuGrade)
	for _, stu := range stuGrade {
		fmt.Printf("ch:%d, en:%d, ma:%d\n", stu.chinese, stu.english, stu.math)

	}
}
