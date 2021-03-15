package redis

import (
	"github.com/garyburd/redigo/redis"
	"gotest.tools/assert"
	"testing"
)

func TestGetRedisClient(t *testing.T) {
	client := GetRedisClient()
	client.Do("set", "ankang", "21")
	reply, _ := redis.String(client.Do("get", "ankang"))
	assert.Equal(t, reply, "21")
}