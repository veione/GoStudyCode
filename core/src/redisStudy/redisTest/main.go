package main

import (
	"fmt"
	"strconv"
	"testpro/redisStudy/logicredis"
)

func main() {
	res := logicredis.GetRedisClient().Get("wtq").Val()
	id, err := strconv.Atoi(res)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
}
