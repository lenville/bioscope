package main

import "fmt"

func main() {
	prtThem(1, 2, 3, 4, 5, 6, 7, 8, 9)
	prtThem(1, 5, 9)
}

func prtThem(numbers ...int) {
	for _, value := range numbers {
		fmt.Print(value, " ")
	}
	fmt.Print("\n")
}
