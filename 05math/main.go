package main

import (
	"crypto/sha256"
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Math operations
	fmt.Println("Square root of 16:", math.Sqrt(16)) // 4
	fmt.Println("Power 2^3:", math.Pow(2, 3))        // 8
	fmt.Println("Absolute of -5:", math.Abs(-5))     // 5
	fmt.Println("Ceil of 3.2:", math.Ceil(3.2))      // 4
	fmt.Println("Floor of 3.8:", math.Floor(3.8))    // 3
	fmt.Println("Pi:", math.Pi)

	// Random Numbers using math
	rand.Seed(time.Now().UnixNano()) // seed to avoid same numbers each run
	fmt.Println("Random int:", rand.Int())
	fmt.Println("Random intn (0-9):", rand.Intn(10))
	fmt.Println("Random float:", rand.Float64())

	// Random number using crypto
	data := "Hello Go"
	hash := sha256.Sum256([]byte(data))
	fmt.Printf("SHA256: %x\n", hash)
}
