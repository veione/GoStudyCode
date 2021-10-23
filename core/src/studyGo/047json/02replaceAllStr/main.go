package main

import (
	"fmt"
	"strings"
)

func main()  {
	old := `狸\u001d\u0025`
	if strings.Contains(old, `\`) {
		newStr := strings.ReplaceAll(old,`\`, `\\`)
		fmt.Printf("***********************登录    昵称：%v",newStr)
	}
}
