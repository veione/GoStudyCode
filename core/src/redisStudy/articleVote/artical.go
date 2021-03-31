package articleVote

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	ONE_WEAK_IN_SCORE = 7 * 86400
	VOTE_SCORE = 432
)

func articleVote(conn *redis.Client, userId, articleId string){
	cutoff := time.Now().Unix() - ONE_WEAK_IN_SCORE
	if conn.ZScore("time", articleId).Val() < float64(cutoff){
		return
	}
	if conn.SAdd("voted:"+articleId, userId).Val() > 0{
		conn.ZIncrBy("score", float64(VOTE_SCORE), articleId)
		conn.HIncrBy("articleId","votes", 1)
	}
}

func articlePost(conn *redis.Client, userId, title, link string) string{
	articleId := string(conn.Incr("article:").Val())
	voted := "voted:" + articleId
	conn.SAdd(voted, userId)
	conn.Expire(voted,time.Hour*24*8*7)

	article := make(map[string]interface{})
	article["title"] = title
	article["link"] = link
	article["poster"] = userId
	article["time"] = time.Now().Unix()
	article["votes"] = 1
	conn.HMSet(articleId, article)

	conn.ZAdd("score", redis.Z{float64(time.Now().Unix()+VOTE_SCORE), articleId})
	conn.ZAdd("time", redis.Z{float64(time.Now().Unix()), articleId})

	return articleId
}

func getArticles(conn *redis.Client, page int, order string){

}