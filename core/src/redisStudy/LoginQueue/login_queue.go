package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"math"
	"redisstudy/logicredis"
	"time"
)

const (
	maxRetryWaitTime   int     = 30 // 最大的重试时间间隔
	waitTimeBitLen     int     = 20 // 累计等待时间计数占位数
	waitTimeLimit      int64   = 10 // 等待重试最大超时间隔
	cfgLoginWindowSize int64   = 10
	LoginTimeSlice     int64   = 5    // 时间片大小
	EnableLoginQueue   bool    = true //是否使用登录排队功能
	LoginRetryRatio    float64 = 10   // 登录重试时间系数
)

var (
	loginQueueLength    int64 = 0       // 排队队列长度
	maxLoginQueueLength int64 = 1000000 // 排队的最大排队长度 100w
)

// 排队队列的键名
func getLoginQueueKeyName() string {
	return "LoginQueue"
}

func getLoginWindowKeyName(rankTime int64) string {
	return fmt.Sprintf("LoginWindow-%v", rankTime)
}

// 登录的时间片
func getNLoginTimeSlice() int64 {
	cfgTimeSlice := int64(LoginTimeSlice)
	if cfgTimeSlice > 60 {
		cfgTimeSlice = 60
	}
	return time.Now().Unix() / cfgTimeSlice
}

// 排队队列的长度
func getLoginQueueLength() int64 {
	return logicredis.GetRedisClient().ZCard(getLoginQueueKeyName()).Val()
}

// 返回指定id的排名,分数
func getLoginQueueRank(dev string) (int64, int64, int64, bool) {

	var (
		curRank     int64
		curScore    int64
		curWaitTime int64
	)

	loginQueueKeyName := getLoginQueueKeyName()
	idxRet := logicredis.GetRedisClient().ZRank(loginQueueKeyName, dev)
	if idxRet.Err() == redis.Nil { // 没有这个元素则
		return curRank, curScore, curWaitTime, false
	} else if idxRet.Err() != nil { // redis出错了不放行
		return math.MaxInt64, 0, 0, true
	} else {
		curRank = idxRet.Val() + 1
	}

	scoreRet := logicredis.GetRedisClient().ZScore(loginQueueKeyName, dev)
	if scoreRet.Err() == redis.Nil {
		return curRank, curScore, curWaitTime, false
	} else if scoreRet.Err() != nil { // redis出错了不放行
		return math.MaxInt64, 0, 0, true
	} else {
		curScore, curWaitTime = getLoginQueueScore(int64(scoreRet.Val()))
	}

	return curRank, curScore, curWaitTime, true
}

// 生成排队分数(类似于排队取号,生成的值就是取的排队号)
// 当前以秒级的时间戳作为排队号
// 排队分数分为2段|score(时间戳)|waitTime(等待时间)|
// score+waitTime(累计的等待时间)=下一次客户端过来重试的时间戳
func genLoginQueueScore(score, lastWaitTime, waitTime int64) int64 {
	nowTime := time.Now().Unix()
	if score == 0 {
		score = nowTime
	}

	// 矫正过去累计的时间(当前时间减去进入排行榜的时间表示玩家已经等待的时间)
	if lastWaitTime != 0 {
		lastWaitTime = nowTime - score
	}

	// 高位表示时间戳，分数大小始终以时间戳排序
	// 地位表示等待时间，时间戳相同时，等待时间越大排名越靠后
	return (score << waitTimeBitLen) | (lastWaitTime + waitTime)
}

// 返回实际的排行榜分数(时间戳)和累计等待(重试)的时间
func getLoginQueueScore(score int64) (int64, int64) {
	var (
		realScore    int64
		waitTime     int64
		waitTimeMask int64 = int64(1<<waitTimeBitLen - 1)
	)
	realScore = score >> waitTimeBitLen
	waitTime = score & waitTimeMask
	return realScore, waitTime
}

// 进入排队队列
func loginQueueEnqueue(dev string) (int64, int64, int64) {

	var (
		curRank     int64
		curScore    int64
		curWaitTime int64
	)

	rankScore := genLoginQueueScore(0, 0, int64(maxRetryWaitTime))
	member := redis.Z{
		Score:  float64(rankScore),
		Member: dev,
	}
	loginQueueKeyName := getLoginQueueKeyName()
	logicredis.GetRedisClient().ZAdd(loginQueueKeyName, member)

	curRank, curScore, curWaitTime, ok := getLoginQueueRank(dev)
	if ok {
		return curRank, curScore, curWaitTime
	} else {
		return math.MaxInt64, math.MaxInt64, 0
	}
}

func loginQueueEnqueueWithWaitTime(dev string, curScore, lastWaitTime, waitTime int64) {

	rankScore := genLoginQueueScore(curScore, lastWaitTime, waitTime)

	member := redis.Z{
		Score:  float64(rankScore),
		Member: dev,
	}

	loginQueueKeyName := getLoginQueueKeyName()

	logicredis.GetRedisClient().ZAdd(loginQueueKeyName, member)
}

// 退出排队队列
func loginQueueDequeue(dev string) {
	loginQueueKeyName := getLoginQueueKeyName()
	logicredis.GetRedisClient().ZRem(loginQueueKeyName, dev)
}

// 获取当前
func getLoginWindowSize(rankTime int64) int64 {
	windowKey := getLoginWindowKeyName(rankTime)
	windowSize, err := logicredis.GetRedisClient().Get(windowKey).Int64()
	if err != nil && err != redis.Nil {
		return 0
	}

	windowSize = cfgLoginWindowSize - windowSize

	if windowSize < 0 {
		windowSize = 0
	}

	return windowSize
}

// 占用一个登录窗口席位
func occupyingALoginWindowSeat(rankTime int64) bool {
	windowKey := getLoginWindowKeyName(rankTime)
	curValue := logicredis.GetRedisClient().Incr(windowKey)

	if curValue.Err() != nil {
		return false
	}

	cfgTimeSlice := int64(LoginTimeSlice)
	if cfgTimeSlice > 60 {
		cfgTimeSlice = 60
	}

	logicredis.GetRedisClient().Expire(windowKey, time.Duration(2*cfgTimeSlice)*time.Second)

	if curValue.Val() > cfgLoginWindowSize {
		return false
	} else {
		return true
	}
}

func checkLoginWaitLevel(dev string) (int, int) {

	if EnableLoginQueue {
		return 0, 0
	}

	var waitLevel int

	rankTime := getNLoginTimeSlice()

	cfgTimeSlice := int64(LoginTimeSlice)
	if cfgTimeSlice > 60 {
		cfgTimeSlice = 60
	}

	waitLevelRatio := 60 / cfgTimeSlice

	curRank, curScore, curWaitTime, ok := getLoginQueueRank(dev)
	if !ok {
		// 当前排队的人数已经超过队列大小了
		if loginQueueLength > maxLoginQueueLength {
			waitRankTime := loginQueueLength / cfgLoginWindowSize
			waitLevel = int(waitRankTime/waitLevelRatio + 1)
			return waitLevel, maxRetryWaitTime
		}
		curRank, curScore, curWaitTime = loginQueueEnqueue(dev)
	}

	// 当前剩余的窗口数
	loginWindowSize := getLoginWindowSize(rankTime)

	if loginWindowSize <= 0 { // 没有窗口了
		waitRankTime := curRank / cfgLoginWindowSize
		waitLevel = int(waitRankTime/waitLevelRatio + 1)
	} else if curRank <= loginWindowSize { // 还有窗口则尝试一下
		if occupyingALoginWindowSeat(rankTime) { // 抢一个名额
			waitLevel = 0
		} else {
			waitRankTime := curRank / cfgLoginWindowSize
			waitLevel = int(waitRankTime/waitLevelRatio + 1)
		}
	} else {
		waitRankTime := curRank / cfgLoginWindowSize
		waitLevel = int(waitRankTime/waitLevelRatio + 1)
	}

	// 等待等级越大等待重试时间越长
	waitRatio := math.Ceil(float64(curRank) / float64(cfgLoginWindowSize*waitLevelRatio))

	cfgLoginRetryRatio := LoginRetryRatio

	waitTime := int(math.Ceil(cfgLoginRetryRatio * waitRatio))

	if waitTime > maxRetryWaitTime {
		waitTime = maxRetryWaitTime
	}

	if waitLevel == 0 {
		loginQueueDequeue(dev)
	}

	if waitLevel > 0 {
		loginQueueEnqueueWithWaitTime(dev, curScore, curWaitTime, int64(waitTime)) // 将排队分数设置为下一次更新时的时间
	}

	return waitLevel, waitTime
}

//////////////////login进程有一个master进程负责发放登录名额和清除过期排队数据//////////////////////////////////////

func logLoginQueueSizeTimer() {
	loginQueueLength = getLoginQueueLength()
}

func clearLoginQueueTimeoutMemberTimer() {
	clearLoginQueueTimeoutMember()
}

// 清除排队队列中过期的成员,每一个时间窗口检查top 2个时间窗口的数据
func clearLoginQueueTimeoutMember() {

	//if ServerInstance().processIdx != 1 {
	//	return
	//}

	loginQueueKeyName := getLoginQueueKeyName()
	topData := logicredis.GetRedisClient().ZRangeWithScores(loginQueueKeyName, 0, cfgLoginWindowSize).Val()
	if len(topData) <= 0 {
		return
	}

	nowTime := time.Now().Unix()
	clearDevs := make([]interface{}, 0, len(topData))
	for _, data := range topData {
		dev, ok := data.Member.(string)
		if !ok {
			continue
		}
		rankTime, waitTime := getLoginQueueScore(int64(data.Score))
		// retryTime表示预测的客户端来重试的时间戳
		retryTime := rankTime + waitTime
		if retryTime+waitTimeLimit < nowTime {
			clearDevs = append(clearDevs, dev)
		}
	}

	clearDevsLen := len(clearDevs)
	if clearDevsLen > 0 {
		logicredis.GetRedisClient().ZRem(loginQueueKeyName, clearDevs...)
		fmt.Printf("[login] 队列清除离开玩家:%v \n", clearDevsLen)
	}
}
