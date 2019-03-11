package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

const (
	UA = "Golang Downloader from Kejibo.com"
)

func main() {
	f, err := os.OpenFile("./file.exe", os.O_RDWR, 0666) //其实这里的 O_RDWR应该是 O_RDWR|O_CREATE，也就是文件不存在的情况下就建一个空文件，但是因为windows下还有BUG，如果使用这个O_CREATE，就会直接清空文件，所以这里就不用了这个标志，你自己事先建立好文件。
	if err != nil {
		panic(err)
	}
	stat, err := f.Stat() //获取文件状态
	if err != nil {
		panic(err)
	}
	f.Seek(stat.Size(), 0) //把文件指针指到文件末，当然你说为何不直接用 O_APPEND 模式打开，没错是可以。我这里只是试验。
	var req http.Request
	req.Method = "GET"
	req.UserAgent() = UA
	req.Close = true

	if err != nil {
		panic(err)
	}
	header := http.Header{}
	header.Set("Range", "bytes="+strconv.FormatInt(stat.Size(), 10)+"-")
	req.Header = header
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		panic(err)
	}
	written, err := io.Copy(f, resp.Body)
	if err != nil {
		panic(err)
	}
	println("written: ", written)
}
