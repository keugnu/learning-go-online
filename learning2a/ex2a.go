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

	fmt.Printf("Average: %v", avg(arr[:]))
}

func avg(arr []float64) (a float64) {
	var sum float64
	for _, num := range arr {
		sum += num
	}

	a = sum / float64(len(arr))
	return
}
