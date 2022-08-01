package myredis

import (
	"sync"
	"time"

	"github.com/go-redis/redis"
)

/**
 * @Author  jackie.lqj
 * @Date  2022/5/18 14:12
 * @Description
 */

var (
	Client *redis.Client
	once   sync.Once
)

const (
	RedisAddr = "127.0.0.1:6379"

	RedisDialTimeout  = time.Duration(50) * time.Millisecond
	RedisReadTimeout  = time.Duration(100) * time.Millisecond
	RedisWriteTimeout = time.Duration(100) * time.Millisecond
)

func InitializeRedisInstance() {
	once.Do(func() {
		opts := &redis.Options{
			Addr:         RedisAddr,
			DialTimeout:  RedisDialTimeout,
			ReadTimeout:  RedisReadTimeout,
			WriteTimeout: RedisWriteTimeout,
		}
		Client = redis.NewClient(opts)
	})

	Client.Pipeline()
}
