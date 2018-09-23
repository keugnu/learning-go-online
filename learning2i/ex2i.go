package main

import (
	"fmt"
)

var ptr int

func main() {
	var s [100]int
	stack := s[:5]
	var val int
	ptr = -1

	for i := 0; i < 5; i++ {
		stack = push(stack, i+1)
	}

	for _, n := range stack {
		fmt.Printf("%v ", n)
	}

	for i := 0; i < 3; i++ {
		val = pop(stack)
		fmt.Printf("\npop: %v", val)
	}

	stack = push(stack, 11, 12)
	print("\n")

	for i := 0; i < ptr+1; i++ {
		fmt.Printf("%v ", s[i])
	}

}

func push(s []int, n ...int) []int {
	for _, x := range n {
		ptr++
		s[ptr] = x
	}
	return s
}

func pop(s []int) (val int) {
	val = s[ptr]
	ptr--
	return
}
