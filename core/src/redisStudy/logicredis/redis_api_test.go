package logicredis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestCmdGet(t *testing.T) {
	info := GetRedisClient().Get("tybfz2pj")
   if info.Err() != nil && info.Err() == redis.Nil {
	   println("没有该键", info.Err().Error())
   }
   println("成功 ",info.Val())
}

func TestCmdSetNx(t *testing.T) {
	info := GetRedisClient().SetNX("1", "wer", 0)
	if info.Err() != nil {
		fmt.Println(info.Err().Error())
	}
	fmt.Println("值", info.Val())
}