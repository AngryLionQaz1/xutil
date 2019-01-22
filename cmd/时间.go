package main

import (
	"fmt"
	"math"
	"time"
)

const base_format = "2006-01-02 15:04:05"

func main() {

	t1 := time.Date(2018, 1, 10, 0, 0, 1, 100, time.Local)
	t2 := time.Date(2018, 2, 9, 23, 59, 22, 100, time.Local)
	fmt.Println(timeSub(t1, t2))
	s := t1.Format(base_format)
	s_time, _ := time.Parse(base_format, s)

	fmt.Println(s, "=============之前的论剑时间小时是")
	fmt.Println(s_time.Weekday())
}
func timeSub(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int64(math.Abs(float64(t1.Sub(t2).Hours() / 24)))
}
