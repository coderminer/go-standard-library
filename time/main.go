package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))
	fmt.Println(now.Format("2006年01月02日 15:04:05"))
	fmt.Println(now.Format(time.RFC3339))

	t1, err := time.Parse("2006-01-02 15:04:05", "2018-09-21 10:54:11")
	t2, err := time.Parse("2006/01/02 15:04:05", "2018/09/21 10:54:59")
	t3, err := time.Parse("2006年01月02日 15:04:05", "2018年09月21日 10:54:59")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	sec := now.Unix()      //秒
	nsec := now.UnixNano() //纳秒
	fmt.Println(sec)
	fmt.Println(nsec)

	t := time.Unix(sec, 0)
	fmt.Println(t)
	fmt.Println(t.Format("2006-01-02 15:04:05"))

	year, month, day := now.Date()
	fmt.Println(year, month, day)
	fmt.Println(year, int(month), day)
	fmt.Printf("year:%d month:%d day:%d\n", year, month, day)

	hour, minute, second := now.Clock()
	fmt.Println(hour, minute, second)

	weekday := now.Weekday()
	fmt.Println(weekday)      // Friday
	fmt.Println(int(weekday)) //5

	days := now.YearDay()
	fmt.Println(days)

	date1 := time.Date(2017, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	date2 := time.Date(2018, time.Month(1), 1, 0, 0, 0, 0, time.UTC)
	hours := date2.Sub(date1).Hours()
	fmt.Println(hours)
	between_days := hours / 24
	fmt.Println(between_days)

	month_days := time.Date(now.Year(), now.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	fmt.Println(month_days)

	duration := time.Since(start)
	fmt.Println(duration)
	fmt.Println(duration.Nanoseconds())
}
