package main

import (
	"fmt"
)

func main() {
	i := 0
loop:
	fmt.Println("Iteration:", i)
	if i < 9 {
		i++
		goto loop
	}
}
