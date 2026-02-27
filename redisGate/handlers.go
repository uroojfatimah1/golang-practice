package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Response struct {
	Source string   `json:"source"`
	Data   []string `json:"data"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	key := "products"

	cacheData, found := GetCache(key)
	if found {
		w.Write(cacheData)
		rdb.Incr(ctx, "cache_hits")
		LogRequest(r, "INFO", "Cache hit for /products")
		return
	}

	// Simulate DB delay
	time.Sleep(3 * time.Second)

	response := Response{
		Source: "database",
		Data:   []string{"Laptop", "Phone", "Tablet"},
	}

	jsonData, _ := json.Marshal(response)
	SetCache(key, jsonData)
	rdb.Incr(ctx, "cache_misses")
	LogRequest(r, "INFO", "Cache miss for /products")
	w.Write(jsonData)
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	hits, _ := rdb.Get(ctx, "cache_hits").Result()
	misses, _ := rdb.Get(ctx, "cache_misses").Result()

	data := map[string]string{
		"cache_hits":   hits,
		"cache_misses": misses,
	}

	LogRequest(r, "INFO", "Stats requested")
	json.NewEncoder(w).Encode(data)
}
