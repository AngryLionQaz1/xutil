package snow

import (
	"fmt"
	"math"
	"time"
)

const BASE_FORMAT = "2006-01-02"

func main() {

	//t1 := time.Date(2019, 1, 6, 0, 0, 1, 100, time.Local)
	//t2 := time.Date(2019, 1, 20, 23, 59, 22, 100, time.Local)
	//fmt.Println(workerDay("2019-01-06", "2019-01-20"))
	fmt.Println(workerDate("2019-01-06", 10))
	//s := t1.Format(BASE_FORMAT)
	//s_time, _ := time.Parse(BASE_FORMAT, s)
	//
	//fmt.Println(s, "=============之前的论剑时间小时是")
	//fmt.Println(t2.Weekday())
	//fmt.Println(s_time.Weekday())
}

func workerDate(s string, count int64) time.Time {
	start, _ := time.Parse(BASE_FORMAT, s)
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local)
	var d int64
	var end string
	var endDay time.Time
	for d = 0; d < count; d++ {
		endDay = start.AddDate(0, 0, int(d))
		end = endDay.Weekday().String()
		if "Sunday" == end || "Saturday" == end {
			count++
			continue
		}
	}
	return endDay
}

/**
  根据日期计算工作日
*/
func workerDay(s, e string) int64 {
	start, _ := time.Parse(BASE_FORMAT, s)
	end, _ := time.Parse(BASE_FORMAT, e)
	var count int64
	var d int64
	start = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, time.Local)
	end = time.Date(end.Year(), end.Month(), end.Day(), 0, 0, 0, 0, time.Local)
	day := int64(math.Abs(float64(start.Sub(end).Hours() / 24)))
	for d = 0; d <= day; d++ {
		s := start.AddDate(0, 0, int(d)).Weekday().String()
		if "Sunday" == s || "Saturday" == s {
			continue
		}
		count++
	}
	return count
}

func timeSub(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int64(math.Abs(float64(t1.Sub(t2).Hours() / 24)))
}
