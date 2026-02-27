package main

import (
	"net"
	"net/http"
	"time"
)

func RateLimiter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ip, _, _ := net.SplitHostPort(r.RemoteAddr)
		key := "rate:" + ip
		count, _ := rdb.Incr(ctx, key).Result()
		if count == 1 {
			rdb.Expire(ctx, key, time.Hour).Result()
		}
		if count > 10 {
			w.WriteHeader(429)
			w.Write([]byte("Rate limit exceeded"))

			return
		}
		next(w, r)
	}
}
