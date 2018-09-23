package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [100]int

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(500)
	}

	fmt.Printf("Maximum: %v", max(arr[0:rand.Intn(100)]))
}

func max(arr []int) (max int) {
	for _, n := range arr {
		if n > max {
			max = n
		}
	}
	return
}
