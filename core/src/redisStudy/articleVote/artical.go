package articleVote

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	ONE_WEAK_IN_SCORE = 7 * 86400
	VOTE_SCORE = 432
)

func articleVote(conn *redis.Client, user, article string){
	cutoff := time.Now().Unix() - ONE_WEAK_IN_SCORE
	if conn.ZScore("time", article).Val() < float64(cutoff){
		return
	}
	article_id
}
