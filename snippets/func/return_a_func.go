package main

import (
	"fmt"
)

func plusX(x int) func(int) int {
	return func(n int) int {
		return n + x
	}
}

func main() {
	p := plusX(10)
	fmt.Printf("%v\n", p(10))
}
