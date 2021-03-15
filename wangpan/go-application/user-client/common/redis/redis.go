package redis

import (
	"github.com/garyburd/redigo/redis"
	"sync"
	"user-client/common"
)

var (
	log  = common.GetLogger()
	coon redis.Conn
	once sync.Once
)

func GetRedisClient() redis.Conn {
	once.Do(func() {
		loadRedisConf()
	})

	return coon
}

func loadRedisConf() {
	var err error
	coon, err = redis.Dial("tcp", "112.74.166.230:6379")
	if err != nil {
		log.Errorf("redis conn error: %v", err)
		panic(err)
	}

	log.Info("redis load success")
}

func ReleaseRedis() {
	coon.Close()
}
