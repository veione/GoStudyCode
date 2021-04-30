package logicredis

import (
	"github.com/go-redis/redis"
	"testing"
)

func TestRedisGet(t *testing.T) {
	info  := GetRedisClient().Get("tybfz2pj")
   if info.Err() != nil && info.Err() == redis.Nil {
	   println("没有该键", info.Err().Error())
   }
   println("成功 ",info.Val())
}
