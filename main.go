package main

import (
	"fmt"
	"qt"
	"qt/qtDB"
	// _ "qt/qtTime"
	"strings"
)

func main() {
	list := []string{
		"usGOOG",    // Google
		"usDJI",     // Dow Jones
		"usAAPL",    // Apple
		"usAMZN",    // Amazon
		"usATVI",    // 动视暴雪啊
		"usBABA",    // Alibaba
		"usAXP",     // 运通
		"usBIDU",    // 百度
		"r_hk00700", // Tencent
		"r_hkHSI",   // Hang Seng
		"sz000001",
		"sz000002",
		"sz000003",
		"sz000004",
		"sz000005",
		"sz000006",
		"sz000007",
		"sz000008",
		"sz000009",
		"sz000010",
		"sz000011",
		"sz000012",
		"sz000013",
		"sz000014",
		"sz000015",
		"sz000016",
		"sz000017",
		"sz000018",
		"sz000019",
		"sz000020",
		"sz000021",
		"sz000022",
		"sz000023",
		"sz000024",
		"sz000025",
		"sz000026",
		"sz000027",
		"sz000028",
		"sz000029",
		"sz000030",
		"sh000001",
		"sh000002",
		"sh000003",
		"sh000004",
		"sh000005",
		"sh000006",
		"sh000007",
		"sh000008",
		"sh000009",
		"sh000010",
		"sh000011",
		"sh000012",
		"sh000013",
		"sh000015",
		"sh000016",
		"sh000017",
		"sh000018",
		"sh000019",
		"sh000020",
		"sh000021",
		"sh000022",
		"sh000023",
		"sh000025",
		"sh000026",
		"sh000027",
		"sh000028",
		"sh000029",
		"sh000030",
	}

	compareBuf := make(map[string]string, len(list))
	cacheBuf1 := make([]string, 0)
	cacheBuf2 := make([]string, 0)
	cacheMark := 1
	const (
		MAXBUFF = 1000
	)

	qtDB.Create(list)
	var counter int = 0

	for {
		counter = counter + 1
		fmt.Printf("%d\n", counter)
		for _, n := range list {
			qtString := qt.GetQtData(n)
			qtArray := strings.Split(qtString, "~")
			time := qtArray[30]
			if compareBuf[n] == time {
				continue
			}
			compareBuf[n] = time
			if cacheMark > 0 {
				cacheBuf1 = append(cacheBuf1, qtString)
				fmt.Printf("Buf1 -> %d - %s  %s", len(cacheBuf1), n, time)
				if len(cacheBuf1) >= MAXBUFF {
					cacheMark = -1
					go qtDB.Insert(cacheBuf1)
					cacheBuf1 = make([]string, 0)
				}
			} else {
				cacheBuf2 = append(cacheBuf2, qtString)
				fmt.Printf("Buf2 -> %d - %s  %s", len(cacheBuf2), n, time)
				if len(cacheBuf2) >= MAXBUFF {
					cacheMark = 1
					go qtDB.Insert(cacheBuf2)
					cacheBuf2 = make([]string, 0)
				}

			}

		}
	}
}
