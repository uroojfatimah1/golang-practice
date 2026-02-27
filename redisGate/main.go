package main

import (
	"fmt"
	"net/http"
)

func main() {
	RedisInit()
	http.HandleFunc("/products", RateLimiter(ProductsHandler))
	http.HandleFunc("/stats", StatsHandler)

	fmt.Println("Server starting on port 8080")
	http.ListenAndServe(":8080", nil)
}
