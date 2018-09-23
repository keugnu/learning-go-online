package main

import (
	"fmt"
)

func main() {
	sample("a", "list", "of", "strings", "for", "a", "variadic", "function")
}

func sample(args ...string) {
	for _, a := range args {
		fmt.Printf("%s ", a)
	}
}
