package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var arr [10]int
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(500)
		fmt.Printf("%v ", arr[i])
	}

	doublerMap(double, arr[:])
	println()

	for _, n := range arr {
		fmt.Printf("%v ", n)
	}
}

func doublerMap(f func(int) int, arr []int) []int {
	for i, n := range arr {
		arr[i] = f(n)
	}
	return arr
}

func double(n int) int {
	return n * 2
}
