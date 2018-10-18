package main

import (
	"strconv"
	"time"
	"fmt"
)

var week time.Duration

func main()  {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(), t.Month(), t.Year())
	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 60 * 60 * 24 * 7 * 1e9
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)
	// formatting times:
	fmt.Println(t.Format(time.RFC822)) // 21 Dec 11 0852 UTC
	fmt.Println(t.Format(time.ANSIC)) // Wed Dec 21 08:56:34 2011
	fmt.Println(t.Format("02 Jan 2006 15:04")) // 21 Dec 2011 08:52
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
	// Wed Dec 21 08:52:14 +0000 UTC 2011 => 20111221
	fmt.Println()
	fmt.Println()
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	fmt.Println(firstOfMonth)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	fmt.Println(firstOfMonth.Unix())
	fmt.Println(lastOfMonth.Unix())
	fmt.Println()

	toBeCharge := "2018-10-01 00:00:00"
	timeLayout := "2006-01-02 03:04:05"                             //转化所需模板
	loc, _ := time.LoadLocation("Asia/Shanghai")                            //重要：获取时区
	startOfMonth, _ := time.ParseInLocation(timeLayout, toBeCharge, loc)
	endOfMonth := startOfMonth.AddDate(0, 1, -1)
	fmt.Println(startOfMonth, endOfMonth)
	fmt.Println()
	todayStr := now.Format("2006-01-02")
	fmt.Println(todayStr)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	fmt.Println(today.Unix())
	fmt.Println(today.AddDate(0, 0, 1).Unix())
	fmt.Println(3600 * 24)

	timeToday, _ := time.Parse("2006-01-02", todayStr)
	nextDay := timeToday.AddDate(0, 0, 1)
	fmt.Println(timeToday, nextDay)
	all := 3600 * 12 + 200
	hours := all / 3600
	minute := all % 3600 / 60
	fmt.Println(hours, minute)
	fmt.Println(strconv.FormatInt(now.Unix(), 10))
}
