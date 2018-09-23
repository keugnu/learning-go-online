package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr [10]int
	for i := 0; i < 10; i++ {
		arr[i] = rand.Int()
	}
	fmt.Print(arr)
}
