package signServer

import (
	"context"
)

// TotalToday 签到总人次, 返回今日第几个签到
func TotalToday() int64 {
	key := getKeyTotalToday()
	result, err := rdb().Incr(context.Background(), key).Result()
	if err != nil {
		panic(err)
	}

	return result
}

// GetTodayTotalNum 获取今日签到人数
func GetTodayTotalNum() int64 {
	key := getKeyTotalToday()
	result, _ := rdb().Get(context.Background(), key).Int64()

	return result
}

// GetYesterdayNum 获取昨日签到人数
func GetYesterdayNum() int64 {
	key := getKeyTotalYesterday()
	result, err := rdb().Get(context.Background(), key).Int64()
	if err != nil {
		panic(err)
	}
	return result
}
