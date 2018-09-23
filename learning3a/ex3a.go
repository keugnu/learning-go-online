package main

import (
	"fmt"
	"stack"
)

func main() {
	s := new(stack.Stack)
	s.Push(1)
	s.Push(8)
	for i := 0; i < 2; i++ {
		fmt.Printf("%v", s.Pop())
	}
}
