package main

import (
	"fmt"
	"time"
)

var c chan int
var i int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready!")
	c <- 1
}

func main() {
	i = 0
	c = make(chan int)
	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting, but not too long")
L:
	for {
		select {
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}
}
