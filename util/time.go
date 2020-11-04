package util

import "time"

var loc *time.Location

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

func FormateYMDH() string {
	if loc == nil {
		loc, _ = time.LoadLocation("Asia/Shanghai")
	}
	return time.Now().In(loc).Format("2006010215")
}

func NowHour() int {
	return time.Now().In(loc).Hour()
}

func NowMonth() int {
	return int(time.Now().In(loc).Month())
}

func NowYear() int {
	return time.Now().In(loc).Year()
}

func NowMinute() int {
	return time.Now().In(loc).Minute()
}

func NowSecond() int {
	return time.Now().In(loc).Second()
}

func NowDay() int {
	return time.Now().In(loc).Day()
}

func NowTimeSteamp() int64 {
	return time.Now().In(loc).Unix()
}