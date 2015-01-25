package main

import "fmt"

func Map(f func(int) int, l []int) []int {
	j := make([]int, len(l))
	for key, value := range l {
		j[key] = f(value)
	}
	return j
}

func main() {
	m := []int{1, 3, 4}
	f := func(i int) int {
		return i * i
	}
	fmt.Printf("%v", Map(f, m))
}
