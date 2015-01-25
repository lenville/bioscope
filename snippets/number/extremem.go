package main

import "fmt"

func max(l []int) int {
	max := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] > max {
			max = l[i]
		}
	}
	return max
}

func min(l []int) int {
	min := l[0]
	for i := 1; i < len(l); i++ {
		if l[i] < min {
			min = l[i]
		}
	}
	return min
}

func main() {
	l := []int{0, 1, 2, 3, 4, -5, -6, -7, 8, 9, 10}
	fmt.Println(max(l))
	fmt.Println(min(l))

}
