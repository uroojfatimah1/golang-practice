package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

func RedisInit() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}
