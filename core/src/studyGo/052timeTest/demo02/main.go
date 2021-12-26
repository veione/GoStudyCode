package main

import (
	"fmt"
	"time"
)

// GetDayStartTM ...
func GetDayStartTM(tm int64) int64 {
	var d time.Time
	if tm != 0 {
		d = time.Unix(tm, 0)
	} else {
		d = time.Now()
	}

	if d.Hour() < 5 {
		return time.Date(d.Year(), d.Month(), d.Day(), 5, 0, 0, 0, d.Location()).Unix() - 86400
	} else {
		return time.Date(d.Year(), d.Month(), d.Day(), 5, 0, 0, 0, d.Location()).Unix()
	}

}


// GetNextDayStartTM: 返回当前时间的  下一天5点
func GetNextDayStartTM(tm int64) int64 {
	var d time.Time
	if tm != 0 {
		d = time.Unix(tm, 0)
	} else {
		d = time.Now()
	}
	if d.Hour() < 5 {
		return time.Date(d.Year(), d.Month(), d.Day(), 5, 0, 0, 0, d.Location()).Unix()
	} else {
		return time.Date(d.Year(), d.Month(), d.Day()+1, 5, 0, 0, 0, d.Location()).Unix()
	}
}


// 获取上一个 周1 凌晨5点 时间戳
func GetLastWeekDayUnixTime(nowTime time.Time, targetWeek time.Weekday)int64{
	nowWeekDay := nowTime.Weekday()
	var offset int
	if nowWeekDay >= targetWeek {
		offset = int(nowWeekDay - targetWeek)
	}else {
		offset = int(time.Saturday - targetWeek) + int(nowWeekDay + 1)
	}
	lastTargetWeekTime := nowTime.AddDate(0,0, -offset)
	//res := time.Date(lastTargetWeekTime.Year(), lastTargetWeekTime.Month(), lastTargetWeekTime.Day(), 5,0,0, 0, lastTargetWeekTime.Location()).Unix()
	fmt.Println("***************", lastTargetWeekTime)
	res := lastTargetWeekTime.Unix()
	if lastTargetWeekTime.Hour() < 5 {
		res = GetNextDayStartTM(lastTargetWeekTime.Unix())
	}else {
		res = GetDayStartTM(lastTargetWeekTime.Unix())
	}
	return res
}

func main()	{
	nowTime := time.Date(2021,12,21, 3, 0, 0, 0, time.Now().Location())

	for targetWeek := time.Sunday; targetWeek <= time.Saturday; targetWeek++ {
		res := GetLastWeekDayUnixTime(nowTime, targetWeek)
		tm := time.Unix(res, 0)
		fmt.Println(tm)
	}
}