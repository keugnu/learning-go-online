package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var arr [10]float64

	for i := 0; i < 10; i++ {
		arr[i] = rand.Float64()
	}

	sum := 0.0
	for i := 0; i < 10; i++ {
		sum = sum + arr[i]
	}

	fmt.Printf("Average: %v", sum/float64(len(arr)))
}
