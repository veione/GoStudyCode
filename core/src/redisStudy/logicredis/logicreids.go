package logicredis

import (
	"github.com/go-redis/redis"
	"log"
	"sync"
)

var (
	redisClient *redis.Client
	once        sync.Once
	//url         = "redis://81.68.166.65:6379"
	//url      = "redis://175.97.169.134:6379"
	url = "redis://10.1.1.95:6379/2"

	db       = 2
	poolSize = 10
)

func NewRedisClient(url string, db, poolSize int) (*redis.Client, error) {
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

func GetRedisClient() *redis.Client {
	once.Do(func() {
		var error error
		redisClient, error = NewRedisClient(url, db, poolSize)
		if error != nil {
			log.Println(error.Error())
		}
	})
	return redisClient
}

func RedisClient() *redis.Client {
	return redisClient
}
