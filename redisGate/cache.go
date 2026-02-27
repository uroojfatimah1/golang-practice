package main

import (
	"time"

	"github.com/redis/go-redis/v9"
)

func GetCache(key string) ([]byte, bool) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, false
	}
	return []byte(val), true
}

func SetCache(key string, val []byte) {
	rdb.Set(ctx, key, val, 60*time.Second)
}
