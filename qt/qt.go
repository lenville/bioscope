package qt

import (
	"bytes"
	"github.com/qiniu/iconv"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func getURL(code string) string {
	url := bytes.Buffer{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	url.WriteString("http://qt.gtimg.cn/")
	url.WriteString(strconv.Itoa(r.Int()))
	url.WriteString("q=")
	url.WriteString(code)
	return url.String()
}

func GetQtData(code string) string {
	cd, err := iconv.Open("utf-8", "gbk")
	u, _ := url.Parse(getURL(code))
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
		return "error"
	}
	result, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return "error"
	}
	return cd.ConvString(string(result))
}
