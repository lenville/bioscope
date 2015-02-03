package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	fib(10, ch)
}

func calculate(ch chan int) {
	tmp1 := <-ch
	tmp2 := <-ch
	sum := (tmp1 + tmp2)
	fmt.Printf("%d ", sum)
	ch <- tmp2
	ch <- sum
}

func fib(num int, ch chan int) {
	for i := 0; i < num; i++ {
		if i < 2 {
			ch <- 1
			fmt.Printf("%d ", 1)
			continue
		}
		calculate(ch)
	}
}
