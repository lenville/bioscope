// Package qtTime implements a simple library
// for parsing QT Time to normal Unix time
package qtTime

import (
	"strconv"
	"strings"
	"time"
)

type QtTime time.Time

// HS time format YYYYMMDDhhmmss
func ParseHS(hsTime string) time.Time {
	t, _ := strconv.Atoi(hsTime)

	year := t / 10000000000
	mon := t % 1000000000 / 100000000
	day := t % 10000000 / 1000000
	hour := t % 1000000 / 10000
	min := t % 10000 / 100
	sec := t % 100
	month := time.Month(mon)

	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

// HK & US time format YYYY/MM.DD hh:mm:ss
func ParseHK(hkTime string) time.Time {
	t := strings.Fields(hkTime)

	dates := strings.Split(t[0], "/")
	times := strings.Split(t[1], ":")
	year, _ := strconv.Atoi(dates[0])
	mon, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	hour, _ := strconv.Atoi(times[0])
	min, _ := strconv.Atoi(times[1])
	sec, _ := strconv.Atoi(times[2])

	month := time.Month(mon)

	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

func ParseUS(usTime string) time.Time {
	t := strings.Fields(usTime)

	dates := strings.Split(t[0], "-")
	times := strings.Split(t[1], ":")
	year, _ := strconv.Atoi(dates[0])
	mon, _ := strconv.Atoi(dates[1])
	day, _ := strconv.Atoi(dates[2])
	hour, _ := strconv.Atoi(times[0])
	min, _ := strconv.Atoi(times[1])
	sec, _ := strconv.Atoi(times[2])

	month := time.Month(mon)

	return time.Date(year, month, day, hour, min, sec, 0, time.UTC)
}

func GetDate(t time.Time) string {
	s := []string{strconv.Itoa(t.Year()), strconv.Itoa(int(t.Month())), strconv.Itoa(t.YearDay())}
	return strings.Join(s, "")
}
