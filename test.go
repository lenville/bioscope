package main

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"log"
	"unicode/utf8"
	"github.com/qiniu/iconv"
)


func main() {
	cd, err := iconv.Open("utf-8", "gbk") // convert utf-8 to gbk
	if err != nil {
        fmt.Println("iconv.Open failed!")
        return
    }
    defer cd.Close()

	u, _ := url.Parse("http://qt.gtimg.cn/q=r_hk00700")
    q := u.Query()
    // q.Set("username", "user")
    // q.Set("password", "passwd")
    u.RawQuery = q.Encode()
    res, err := http.Get(u.String());
    if err != nil {
      log.Fatal(err)
      return;
    }
    result, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
      log.Fatal(err)
      return;
    }
    result := cd.ConvString(result)
    fmt.Printf("%s", result)
	// u, _ := url.Parse("http://qt.gtimg.cn/q=r_hk00700")
	// q := u.Query()
	// u.RawQuery =q.Encode()
	// res, err := http.Get(u.String())
	// result, err := ioutil.ReadAll(res.Body)
	// res.Body.Close()
	// fmt.Printf("%s", result)
	// // q := u.Query()
	// // log(u)
	// fmt.Println("Hello, 世界")
}
