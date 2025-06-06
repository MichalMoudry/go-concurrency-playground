package transport

import "github.com/gomodule/redigo/redis"

func NewRedisPool(redisStr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle: 10,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisStr)
		},
	}
}
