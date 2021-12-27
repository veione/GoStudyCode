package main

import "time"

type tokenBucket struct {
	putTokenBucket uint32   // 每秒放入令牌的个数
	refreshTime    int64    // 最后刷新时间
	capTokenBucket uint32	// 令牌桶容量
	curTokenNum    uint32   // 当前剩余令牌数
}

func NewTokenBucket(cap, curTokenNum uint32) *tokenBucket {
	return &tokenBucket{
		capTokenBucket: cap,
		curTokenNum: curTokenNum,
	}
}

func (bucket *tokenBucket) tryAcquire()bool  {
	nowTime := time.Now().Unix()
	addTokenNum := uint32(nowTime - bucket.refreshTime) * bucket.putTokenBucket
	if bucket.curTokenNum + addTokenNum > bucket.capTokenBucket {
		bucket.curTokenNum = bucket.capTokenBucket
	}else {
		bucket.curTokenNum += addTokenNum
	}
	bucket.refreshTime = nowTime
	if bucket.curTokenNum > 0 {
		bucket.curTokenNum --
		return true
	}
	return false
}

