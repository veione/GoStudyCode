package logicredis

import (
	"github.com/go-redis/redis"
	"log"
)

var redisClient *redis.Client

func NewRedisClient(url string, db, poolSize int) (*redis.Client, error){
	opt, err := redis.ParseURL(url)

	if poolSize > 0 {
		opt.PoolSize = poolSize
	}
	if db > 0 {
		opt.DB = db
	}
	if err != nil {
		log.Printf("new redis client url:%v error:%v", url, err)
		return nil, err
	}
	redisClient := redis.NewClient(opt)
	pong, err := redisClient.Ping().Result()

	if err != nil {
		log.Printf("redis client pong:%v error:%v", pong, err)
		return nil, err
	}

	log.Printf("[redis] new redis url:%v pong:%v", url, pong)

	return redisClient, nil
}

func GetRedisClient() *redis.Client{
	return redisClient
}

