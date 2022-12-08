package signServer

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	_db   *gorm.DB
	_rdb  *redis.Client
	_conf *Config
)

const (
	TagTodayTotal = "sign:today" // 今日签到
	TagUser       = "sign:user"  // 签到人
)

func Init(db *gorm.DB, rdb *redis.Client, config *Config) {
	_rdb = rdb
	_db = db

	if config != nil {
		_conf = config
	} else {
		// 由于未传入起始天, 可能导致偏移量过大
		_conf = &Config{}
	}

}

func conf() *Config {
	return _conf
}

func db() *gorm.DB {
	return _db
}

func rdb() *redis.Client {
	return _rdb
}
