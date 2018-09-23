package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var arr [100]int

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(10000)
	}

	bubbleSort(arr[:])
	printResults(arr[:])
}

func bubbleSort(arr []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				temp := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = temp
				swapped = true
			}
		}
	}
}

func printResults(arr []int) {
	for i, n := range arr {
		if i%(len(arr)/5) == 0 {
			println("")
		}
		fmt.Printf("%4d ", n)
	}
	println("\n")
}
