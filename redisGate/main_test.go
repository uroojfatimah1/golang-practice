package main

import (
	"io"
	"net/http"
	"testing"
	"time"
)

const baseURL = "http://localhost:8080"

//
// TEST 1 — Products Endpoint (Cache Miss)
//

func TestProductsCacheMiss(t *testing.T) {

	start := time.Now()

	resp, err := http.Get(baseURL + "/products")

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	duration := time.Since(start)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200 got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	t.Log("Response:", string(body))

	// First call should be slow
	if duration < 2*time.Second {
		t.Error("Expected cache miss (slow response)")
	}

}

//
// TEST 2 — Products Endpoint (Cache Hit)
//

func TestProductsCacheHit(t *testing.T) {

	time.Sleep(1 * time.Second)

	start := time.Now()

	resp, err := http.Get(baseURL + "/products")

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	duration := time.Since(start)

	body, _ := io.ReadAll(resp.Body)

	t.Log("Response:", string(body))

	// Should be fast
	if duration > 1*time.Second {
		t.Error("Expected cache hit (fast response)")
	}

}

//
// TEST 3 — Stats Endpoint
//

func TestStatsEndpoint(t *testing.T) {

	resp, err := http.Get(baseURL + "/stats")

	if err != nil {
		t.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200 got %d", resp.StatusCode)
	}

	body, _ := io.ReadAll(resp.Body)

	t.Log("Stats:", string(body))

}

//
// TEST 4 — Rate Limiting
//

func TestRateLimit(t *testing.T) {

	var lastStatus int

	for i := 0; i < 12; i++ {

		resp, err := http.Get(baseURL + "/products")

		if err != nil {
			t.Fatal(err)
		}

		lastStatus = resp.StatusCode

		resp.Body.Close()

	}

	if lastStatus != 429 {
		t.Error("Expected rate limit 429 error")
	}

}
