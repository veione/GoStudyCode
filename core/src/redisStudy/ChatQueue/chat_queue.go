package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"redisstudy/logicredis"
	"math"
	"time"
)

var (
	cfgCrossSrvWindowSize int64 = 20
	cfgCrossSrvQueueSize  int64 = 100 //
	cfgTimeSlice      int64 = 3  // 每3秒 有一个计数
)


func getNTimeSlice() int64 {
	return time.Now().Unix() / cfgTimeSlice
}


func getCrossSrvQueueKeyName() string {
	return "CrossSrvChat"
}

func getCrossSrvWindowKeyName(rankTime int64) string {
	return fmt.Sprintf("CrossSrvChatWindow-%v", rankTime)
}

// 生成排队分数
// 当前以秒级别的时间戳作为排队号
func getCrossSrvQueueScore() int64 {
	return time.Now().Unix()
}

// 获取指定id的排名
func getCrossSrvQueueRank(uidStr string) (int64, bool) {
	crossSrvQueueKeyName := getCrossSrvQueueKeyName()
	ret := logicredis.GetRedisClient().ZRank(crossSrvQueueKeyName, uidStr)
	if ret.Err() != nil {
		return 0, false
	}
	return ret.Val() + 1, true
}

// 进入排队队列
func crossSrvQueueEnqueue(uidStr string) int64 {
	var curRank int64
	score := getCrossSrvQueueScore()
	member := redis.Z{
		Score:  float64(score),
		Member: uidStr,
	}
	key := getCrossSrvQueueKeyName()
	logicredis.GetRedisClient().ZAdd(key, member)

	ret := logicredis.GetRedisClient().ZRank(key, uidStr)
	if ret.Err() != nil {
		log.Println(ret.Err().Error())
		return math.MaxInt64
	}
	curRank = ret.Val() + 1
	return curRank
}

// 退出排队队列
func crossSrvQueueDequeue(uidStr string) {
	key := getCrossSrvQueueKeyName()
	logicredis.GetRedisClient().ZRem(key, uidStr)
}

// 获取当前剩余窗口数量
func getCrossSrvWindowSize(rankTime int64) int64 {
	windowKey := getCrossSrvWindowKeyName(rankTime)
	windowSize, err := logicredis.GetRedisClient().Get(windowKey).Int64()
	if err != nil && err != redis.Nil {
		return 0
	}
	windowSize = cfgCrossSrvWindowSize - windowSize
	if windowSize < 0 {
		windowSize = 0
	}
	return windowSize
}

// 占据窗口的一个位置
func occupyingACrossSrvWindowSeat(rankTime int64) bool {
	windowKey := getCrossSrvWindowKeyName(rankTime)
	curValue := logicredis.GetRedisClient().Incr(windowKey)
	if curValue.Err() != nil {
		log.Println(curValue.Err().Error())
		return false
	}
	if curValue.Val() > cfgCrossSrvWindowSize {
		return false
	} else {
		return true
	}
}

// 参数： 0：无需排队  1：可以进入队列   2：队列已满
func checkSendCrossSrvChat(uidStr string) int {
	curRank := crossSrvQueueEnqueue(uidStr)

	if curRank > cfgCrossSrvQueueSize {
		crossSrvQueueDequeue(uidStr)
		return 2
	}
	rankTime := getNTimeSlice()
	// 剩余窗口数
	crossSrvWindowSize := getCrossSrvWindowSize(rankTime)
	// 当前排名小于可用窗口大小 表示可以优先进入
	if curRank <= crossSrvWindowSize {
		if occupyingACrossSrvWindowSeat(rankTime) {
			crossSrvQueueDequeue(uidStr)
			return 0
		}
	}
	return 1
}

// 参数：0：重试成功 1: 继续等待  2：已经不在队列中了，不需要重试了
func retryCheckSendCrossSrvChat(uidStr string) int {
	curRank, ok := getCrossSrvQueueRank(uidStr)
	if !ok {
		return 2
	}
	rankTime := getNTimeSlice()
	// 剩余窗口数
	crossSrvWindowSize := getCrossSrvWindowSize(rankTime)
	// 当前排名小于可用窗口大小 表示可以优先进入
	if curRank <= crossSrvWindowSize {
		if occupyingACrossSrvWindowSeat(rankTime) {
			crossSrvQueueDequeue(uidStr)
			return 0
		}
	}
	return 1
}

func clearCrossSrvQueueTimeoutMemberTimer() {
	clearCrossSrvQueueTimeoutMember(2 * 60)
}

func clearCrossSrvQueueTimeoutMember(timeout int64) {
	var (
		minValue string
		maxValue string
	)
	curCrossSrvQueueScore := getCrossSrvQueueScore()
	if timeout == 0 {
		minValue = "0"
		maxValue = fmt.Sprintf("%v", curCrossSrvQueueScore)
	} else {
		timeoutScore := curCrossSrvQueueScore - timeout
		if timeoutScore < 0 {
			return
		}
		minValue = "0"
		maxValue = fmt.Sprintf("%v", timeoutScore)
	}
	crossSrvQueueKeyName := getCrossSrvQueueKeyName()
	logicredis.GetRedisClient().ZRemRangeByScore(crossSrvQueueKeyName, minValue, maxValue)
}
