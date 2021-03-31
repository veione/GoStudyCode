package LoginQueue

import (
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"logic/logicredis"
	"math"
	"time"
)

const (
	LoginTimeSlice = 60
	LoginWindowSize = 1000
	EnableLoginQueue = true
	LoginRetryRatio = 1.0 // 登录重试时间系数
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
	return logicredis.RateLimitRedis().ZCard(getLoginQueueKeyName()).Val()
}

// 返回指定id的排名,分数
func getLoginQueueRank(dev string) (int64, bool) {
	loginQueueKeyName := getLoginQueueKeyName()
	ret1 := logicredis.RateLimitRedis().ZRank(loginQueueKeyName, dev)
	if ret1.Err() != nil {
		return 0, false
	}
	return ret1.Val() + 1, true
}

// 生成排队分数(类似于排队取号,生成的值就是取的排队号)
// 当前以毫秒级的时间戳作为排队号
func genLoginQueueScore() int64 {
	return time.Now().Unix()
}

// 进入排队队列
func loginQueueEnqueue(dev string) int64 {
	var curRank int64
	rankScore := genLoginQueueScore()
	member := redis.Z{
		Score:  float64(rankScore),
		Member: dev,
	}
	loginQueueKeyName := getLoginQueueKeyName()
	logicredis.RateLimitRedis().ZAdd(loginQueueKeyName, member)

	ret := logicredis.RateLimitRedis().ZRank(loginQueueKeyName, dev)
	if ret.Err() != nil && ret.Err() != redis.Nil {
		log.Println("[redis] loginQueueEnqueue error:%v", ret.Err())
		return math.MaxInt64
	}

	curRank = ret.Val() + 1

	return curRank
}

// 退出排队队列
func loginQueueDequeue(dev string) {
	loginQueueKeyName := getLoginQueueKeyName()
	logicredis.RateLimitRedis().ZRem(loginQueueKeyName, dev)
}

// 获取当前
func getLoginWindowSize(rankTime int64) int64 {
	windowKey := getLoginWindowKeyName(rankTime)
	windowSize, err := logicredis.RateLimitRedis().Get(windowKey).Int64()
	if err != nil && err != redis.Nil {
		log.Println("getLoginWindowSize error:%v", err)
		return 0
	}

	cfgLoginWindowSize := int64(LoginWindowSize)

	windowSize = cfgLoginWindowSize - windowSize

	if windowSize < 0 {
		windowSize = 0
	}

	return windowSize
}

// 占用一个登录窗口席位
func occupyingALoginWindowSeat(rankTime int64) bool {
	windowKey := getLoginWindowKeyName(rankTime)
	curValue := logicredis.RateLimitRedis().Incr(windowKey)

	if curValue.Err() != nil {
		log.Println("occupyingALoginWindowSeat error:%v", curValue.Err())
		return false
	}

	cfgTimeSlice := int64(LoginTimeSlice)
	if cfgTimeSlice > 60 {
		cfgTimeSlice = 60
	}

	logicredis.RateLimitRedis().Expire(windowKey, time.Duration(2*cfgTimeSlice)*time.Second)

	cfgLoginWindowSize := int64(LoginWindowSize)

	if curValue.Val() > cfgLoginWindowSize {
		return false
	} else {
		return true
	}
}

func checkLoginWaitLevel(dev string) (int, int) {

	if !EnableLoginQueue {
		return 0, 0
	}

	var waitLevel int

	rankTime := getNLoginTimeSlice()

	curRank, ok := getLoginQueueRank(dev)
	if !ok {
		curRank = loginQueueEnqueue(dev)
	}

	cfgLoginWindowSize := int64(LoginWindowSize)

	// 当前剩余的窗口数
	loginWindowSize := getLoginWindowSize(rankTime)

	cfgTimeSlice := int64(LoginTimeSlice)
	if cfgTimeSlice > 60 {
		cfgTimeSlice = 60
	}

	waitLevelRatio := 60 / cfgTimeSlice

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

	if waitLevel == 0 {
		loginQueueDequeue(dev)
	} else if waitLevel > 10 {
		waitLevel = 10
	}

	// 等待等级越大等待重试时间越长
	waitRatio := math.Ceil(float64(curRank) / float64(cfgLoginWindowSize))

	cfgLoginRetryRatio := LoginRetryRatio

	waitTime := int(math.Ceil(cfgLoginRetryRatio * waitRatio))

	// waitTime := int(math.Ceil(cfgLoginRetryRatio * waitRatio * waitRatio))

	// if waitLevel == 0 {
	// 	logger.Debugf("[login] curRank:%v loginWindowSize:%v waitLevel:%v curTime:%v nowTime:%v", curRank, loginWindowSize, waitLevel, rankTime, time.Now())
	// }
	return waitLevel, waitTime
}

//////////////////login进程有一个master进程负责发放登录名额和清除过期排队数据//////////////////////////////////////
// 清除排队队列中过期的成员,暂定清除1分钟之前的

func logLoginQueueSizeTimer() {
	if ServerInstance().processIdx == 1 {
		loginQueueLength := getLoginQueueLength()
		if loginQueueLength > 0 {
			log.Printf("[login] 当前排队队列长度:%v", loginQueueLength)
		}
	}
}

func clearLoginQueueTimeoutMemberTimer() {
	if ServerInstance().processIdx == 1 {
		clearLoginQueueTimeoutMember(60 * 60)
	}
}

func clearLoginQueueTimeoutMember(timeout int64) {

	if ServerInstance().processIdx == 1 {

		var (
			minValue string
			maxValue string
		)

		if timeout == 0 {
			minValue = "0"
			curLoginQueueScore := genLoginQueueScore()
			maxValue = fmt.Sprintf("%v", curLoginQueueScore)
		} else {
			curLoginQueueScore := genLoginQueueScore()
			timeoutScore := curLoginQueueScore - timeout
			if timeoutScore < 0 {
				return
			}
			minValue = "0"
			maxValue = fmt.Sprintf("%v", timeoutScore)
		}

		loginQueueKeyName := getLoginQueueKeyName()
		logicredis.RateLimitRedis().ZRemRangeByScore(loginQueueKeyName, minValue, maxValue)
	}
}
