package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	r, _ := http.Get("http://baidu.com")
	b, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("%s\n", b)
}
