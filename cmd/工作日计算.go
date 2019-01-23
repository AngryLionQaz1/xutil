package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const BASE_FORMAT = "2006-01-02"

func main() {

	//实例化一个命令行程序
	app := cli.NewApp()
	//程序名称
	app.Name = "workDate"
	//程序的用途描述
	app.Usage = "工作日计算"
	//程序的版本号
	app.Version = "1.0.0"
	//预置变量
	var startDate string
	var endDate string
	var count string

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "startDate, sd",
			Value:       "2019-01-01",
			Usage:       "开始日期",
			Destination: &startDate,
		},
		cli.StringFlag{
			Name:        "endDate, ed",
			Value:       "2019-01-02",
			Usage:       "结束日期",
			Destination: &endDate,
		},
		cli.StringFlag{
			Name:        "count, c",
			Value:       "10",
			Usage:       "工作日",
			Destination: &count,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "wd",
			Aliases: []string{"w"},
			Usage:   "根据日期计算工作日",
			Action: func(c *cli.Context) error {
				fmt.Println(workerDay(startDate, endDate))
				return nil
			},
		},
		{
			Name:    "ws",
			Aliases: []string{"s"},
			Usage:   "根据工作日计算日期",
			Action: func(c *cli.Context) error {
				sx, _ := strconv.ParseInt(count, 10, 64)
				fmt.Println(workerDate(startDate, sx))
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}

/**
  计算结束工作日
*/
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
