package signServer

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	_db  *gorm.DB
	_rdb *redis.Client
)

const (
	TagTodayTotal = "sign:today" // 今日签到
	TagUser       = "sign:user"  // 签到人
)

func Init(db *gorm.DB, rdb *redis.Client) {
	_rdb = rdb
	_db = db
}

func db() *gorm.DB {
	return _db
}

func rdb() *redis.Client {
	return _rdb
}
