package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	in, _ := strconv.Atoi(os.Args[1])
	if in < 3 {
		return
	} else {
		fib(in)
	}
}

func fib(end int) {
	fmt.Printf("1 1 ")
	x1, x2, x3 := 1, 1, 1
	for {
		x3 = x1 + x2
		x1 = x2
		x2 = x3
		if x3 > end {
			break
		}
		fmt.Printf("%d ", x3)
	}
}
