package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strconv"
	"strings"
)

const (
	DOMAIN = "www.w23588.com"
	URL    = "http://www.w23588.com"
	URL2   = "http://www.w33588.com/"
	CHAT   = 20
)

func main() {

	urls := Cawler01()
	url2 := Cawler02(urls)
	WriteToTxt(URL)
	WriteToTxt(URL2)
	for url := range url2 {
		WriteToTxt(url)
	}
}

//写入文件
func WriteToTxt(url string) {
	go func(u string) {
		fmt.Println(u)
		//文件的创建，Create会根据传入的文件名创建文件，默认权限是0666 os.O_APPEND(追加)
		fileObj, err := os.OpenFile("out.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Failed to open the file", err.Error())
			os.Exit(2)
		}
		defer fileObj.Close()
		if _, err := fileObj.WriteString(u + "\n"); err == nil {
			//fmt.Println("Successful writing to the file with os.OpenFile and *File.WriteString method.",content)
		}
	}(url)
}

//解析第二层
func Cawler02(urls <-chan string) <-chan string {
	url := make(chan string, CHAT)
	go func() {
		for u := range urls {
			collector := colly.NewCollector(colly.Async(true))
			collector.OnHTML("div[class=post-content]>a", func(e *colly.HTMLElement) {
				link := e.Attr("onclick")
				url <- JsUrl(link)
			})
			collector.Visit(u)
		}
	}()

	return url
}

//解析第一层
func Cawler01() <-chan string {
	urls := make(chan string, CHAT)
	go func() {
		collector := colly.NewCollector(
			colly.AllowedDomains(DOMAIN),
			colly.AllowURLRevisit(),
			colly.Async(true),
		)
		collector.OnHTML("a[href][class=io]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			//fmt.Printf("Link found: %q -> %s\n", e.Text, link)
			urls <- link
		})
		collector.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting", r.URL.String())
		})
		collector.OnScraped(func(response *colly.Response) {
			close(urls)
		})
		collector.Visit(URL)
	}()

	return urls
}

//处理Unicode url
func JsUrl(url string) string {
	return UnicodeToUrl(url[len("javascript:SN_Go(String.fromCharCode(") : len(url)-2])
}

// 解析 url
func UnicodeToUrl(u string) string {
	var buffer bytes.Buffer
	strs := strings.Split(u, ",")
	for _, v := range strs {
		i, _ := strconv.ParseInt(v, 10, 64)
		buffer.WriteString(string(i))
	}
	return buffer.String()
}
