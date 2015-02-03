package main

import (
	"fmt"
	"os"
)

func main() {
	for i, l := 0, len(os.Args); i < l; i++ {
		fmt.Println(os.Args[i])
	}
}
