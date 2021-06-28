package main

import "fmt"

func convert(s string, numRows int) string {
	ansMap := make(map[int]string)
	res := ""
	godown := true
	curRow := 0
	for i := 0; i < len(s); i++ {
		ansMap[curRow] += string(s[i])
		if godown && curRow+1 < numRows {
			curRow++
		}
		if !godown && curRow-1 >= 0 {
			curRow--
		}
		if curRow == numRows-1 || curRow == 0 {
			godown = !godown
		}
	}
	for i := 0; i < numRows; i++ {
		res += ansMap[i]
	}
	return res
}

func main() {
	s := "abc"
	numRows := 1
	res := convert(s, numRows)
	fmt.Println(res)
}
