package redisTest

import (
	"fmt"
	"redisstudy/logicredis"
	"testing"
)

func TestSetNX(t *testing.T) {
//   if .NonCacheRedis().SetNX(mmidKey, player.playerID, 0).Val() {
}

func TestHGetAll(t *testing.T) {
	res := logicredis.GetRedisClient().HGetAll("234").Val()
	if len(res) == 0 {
		fmt.Printf("len:%v\n", len(res))
	}
	fmt.Printf("%v\n", res)

}