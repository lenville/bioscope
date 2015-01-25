package main

import "fmt"

func bubble(l []int, sort int) []int {

	swapped := false
	for swapped == false {
	Loop:
		for i := 0; i < len(l)-1; i++ {
			if (l[i]-l[i+1])*sort > 0 {
				l[i], l[i+1] = l[i+1], l[i]
				goto Loop
			}
		}
		swapped = true
	}
	return l
}

func main() {
	l := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	l2 := []int{-1, 20, 49, -100, 0, 98, -5}
	fmt.Println(bubble(l, 1))
	fmt.Println(bubble(l, -1))
	fmt.Println(bubble(l2, 1))
	fmt.Println(bubble(l2, -1))
}
