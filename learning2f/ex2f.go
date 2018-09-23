package main

import (
	"fmt"
)

func main() {
	add := adder
	fmt.Printf("%v\n", add(2, 9, 12, 0, 100))
}

func adder(args ...int) (x int) {
	for _, n := range args {
		x += n
	}
	return
}
