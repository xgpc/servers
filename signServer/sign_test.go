package signServer

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"
)

func redisClient() *redis.Client {

	conn := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	err := conn.Ping(context.Background()).Err()
	if err != nil {
		panic("Redis启动失败，" + err.Error())
	}
	//使用0号数据库
	conn.Do(context.Background(), "select", 0)
	return conn
}

func TestUserSign(t *testing.T) {
	t1 := time.Unix(0, 0)
	t2 := time.Unix(1669132800, 0)
	fmt.Println(t1)
	fmt.Println(t2)

	diff := 1669132800.0
	d := int(diff / (3600 * 24) * 10)

	if int(d)%10 > 0 {
		d = d/10 + 1
	}

	fmt.Println(d)
}

func TestBit(t *testing.T) {

	// 连续打卡三天
	//  7 => 0000 0111

	// 连续打卡 5天
	// 127   => 01111 1111

	a := 133
	for i := 0; i < 10; i++ {

		fmt.Println(a & 1)
		a = a >> 1
	}
}

func TestSign(t *testing.T) {
	rdb := redisClient()
	Init(nil, rdb, nil)
	key := getUserKey(7)
	offset := getTodayNum() - conf().NumDaysAgo
	GetTodayTotalNum()
	seta, err := rdb.SetBit(context.Background(), key, offset, 1).Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(seta)
	result, err := rdb.BitField(context.Background(), key, "get", "u16", offset-15).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println(SignUserCheckToday(7))
	a, _ := SignUserGetInfo(7)
	fmt.Println(a)
	fmt.Println(GetTodayTotalNum())
	fmt.Println("输入")
	var i int64
	for i = 0; i < offset+1; i++ {
		r, err := rdb.GetBit(context.Background(), "sign:user:7", i).Result()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Print(r)
	}
}
