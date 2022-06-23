package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

// GetMd5 获取MD5编码后的字符串
func GetMd5(str string) string {
	token := md5.Sum([]byte(str))
	return hex.EncodeToString(token[:])
}

func main() {
	str := "codeC4php97d deviceiosdim_level0gamemezymainPlatformmix_twplatform2014_1629797289657660role_id20005424server0time1654733484version1.1mezycode#$32!$"
	res := GetMd5(str)
	out := strings.ToUpper(res)
	fmt.Println(out)
}
