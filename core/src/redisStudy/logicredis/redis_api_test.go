package logicredis

import (
	"fmt"
	"github.com/go-redis/redis"
	"testing"
)

func TestCmdGet(t *testing.T) {
	info := GetRedisClient().Get("wtqt")
	if info.Err() != nil && info.Err() == redis.Nil {
		println("没有该键", info.Err().Error())
	}
	fmt.Printf("成功 value:%v \n ", info.Val())
	if info.Val() == "1" {
		fmt.Printf("字符串为空")
	}
}

func TestCmdSetNx(t *testing.T) {
	info := GetRedisClient().SetNX("1", "wer", 0)
	if info.Err() != nil {
		fmt.Println(info.Err().Error())
	}
	if info.Val() {
		fmt.Printf("setNX成功 val:%v", info.Val())
	}
	fmt.Println("值", info.Val())
}

func TestCmdSet(t *testing.T) {
	GetRedisClient().Set("wtqt", "1", 0)

}
