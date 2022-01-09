package main

import "fmt"

type MuseumData struct {
	OTime           int64        `bson:"otime"`  // 开园时间
	TotalOperating  uint32       `bson:"toper"`  // 总经营积分
	TotalPraise     uint32       `bson:"ttps"`   // 总点赞数量
	VTodayOperating uint32       `bson:"vtoper"` // 拜访-今日经营积分
	VTodayPraise    uint32       `bson:"vtdps"`  // 拜访-今日点赞数量
	LastOffLineTM   int64        `bson:"lotm"`   // 离线时间
}

func main() {
	ms := new(MuseumData)
	fmt.Println(ms.TotalPraise)
}

