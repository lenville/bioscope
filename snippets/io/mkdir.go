package main

import (
	"os"
)

func main() {
	if _, e := os.Stat("name"); e != nil {
		os.Mkdir("name", 0755)
	} else {
		// error
	}
}
